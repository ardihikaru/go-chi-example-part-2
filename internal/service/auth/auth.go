package auth

import (
	"github.com/ardihikaru/go-chi-example-part-2/pkg/jwtauth"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/logger"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/service/auth"
)

// Service prepares the interfaces related with this auth service
type Service struct {
	*auth.Service
}

type LoginData auth.LoginData
type Token auth.Token

// NewService creates a new auth service
func NewService(log *logger.Logger, jwtExpTime int, tokenAuth *jwtauth.JWTAuth) *Service {
	service := auth.NewService(log, jwtExpTime, tokenAuth)

	return &Service{
		Service: service,
	}
}

// Authorize authorizes user credential
func (svc *Service) Authorize(loginData LoginData) (*Token, error) {
	token, err := svc.Service.Authorize(auth.LoginData(loginData))
	if err != nil {
		return nil, err
	}
	tokenLocalized := Token(*token)

	return &tokenLocalized, err
}
