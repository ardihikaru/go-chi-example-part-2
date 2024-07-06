package session

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ardihikaru/go-chi-example-part-2/pkg/jwtauth"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/logger"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/middleware"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/utils/marshal"
)

const (
	userKey = "user"
)

// Session defines the user raw document to be inserted
type Session struct {
	AccountId string `json:"account_id"`
	UserEmail string `json:"user_email"`
	UserId    string `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Name      string `json:"name"`
}

// ClaimUser defines the claim user
type ClaimUser struct {
	UserId    string `json:"user_id"`
	AccountId string `json:"account_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type Service struct {
	log *logger.Logger
}

// NewService creates a new timeout handler service
func NewService(log *logger.Logger) *Service {
	return &Service{
		log: log,
	}
}

// SessionCtx enriches the request with the captured JWT private claims
func (svc *Service) SessionCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// extracts token from the header
		token, _, err := jwtauth.FromContext(r.Context())
		if err != nil {
			svc.log.Warn(fmt.Sprintf("failed to load access token: %s", err.Error()))
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// token is authenticated, extracts the private claims
		privateClaims := token.PrivateClaims()

		// marshals user claim
		user := ClaimUser{}
		err = marshal.ToStruct(privateClaims[userKey].(map[string]interface{}), &user)
		if err != nil {
			svc.log.Error(fmt.Sprintf("failed to marshal user claim: %s", err.Error()))
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
			return
		}

		// extracts
		session := Session{
			UserId:    user.UserId,
			AccountId: user.AccountId,
			Username:  user.Username,
			Role:      user.Role,
			Name:      user.Name,
		}

		// token is authenticated, enrich token to the request parameter
		ctx := context.WithValue(r.Context(), middleware.SessionKey, session)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
