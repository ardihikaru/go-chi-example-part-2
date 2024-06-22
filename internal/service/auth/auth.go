package auth

import (
	"github.com/ardihikaru/go-chi-example-part-2/pkg/authenticator"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/jwtauth"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/logger"
	"time"
)

const (
	ClaimExpiredInKey = "exp"
	ClaimIssuedAtKey  = "iat"
)

type Token struct {
	AccessToken string  `json:"access_token"`
	ExpiredIn   int64   `json:"expired_in"`
	IssuedAt    int64   `json:"issued_at"`
	Session     Session `json:"session"`
}

// Session defines the user raw document to be inserted
type Session struct {
	AccountId string `json:"account_id"`
	UserEmail string `json:"user_email"`
	Username  string `json:"username"`
	Name      string `json:"name"`
}

// LoginData is the input JSON body captured from the login request
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Service prepares the interfaces related with this auth service
type Service struct {
	log        *logger.Logger
	tokenAuth  *jwtauth.JWTAuth
	jwtExpTime int
}

// NewService creates a new auth service
func NewService(log *logger.Logger, jwtExpTime int, tokenAuth *jwtauth.JWTAuth) *Service {
	return &Service{
		log:        log,
		jwtExpTime: jwtExpTime,
		tokenAuth:  tokenAuth,
	}
}

func (svc *Service) Authorize(loginData LoginData) (*Token, error) {
	// TODO: implements code here

	// builds the JWT claim options
	durationIn := time.Duration(svc.jwtExpTime) * time.Second
	jwtClaims := authenticator.JWTClaims{
		ClaimExpiredInKey: jwtauth.ExpireIn(durationIn),
		ClaimIssuedAtKey:  jwtauth.EpochNow(),

		// TODO: adds other claims
	}

	// begins to create the access token
	accessToken := authenticator.CreateAccessToken(svc.tokenAuth, jwtClaims)

	token := &Token{
		AccessToken: accessToken,
		ExpiredIn:   jwtauth.ExpireIn(durationIn),
		IssuedAt:    jwtauth.EpochNow(),
		Session: Session{
			AccountId: "id123",
			UserEmail: "user@email.com",
			Username:  "user",
			Name:      "Ini Budi",
		},
	}

	return token, nil
}
