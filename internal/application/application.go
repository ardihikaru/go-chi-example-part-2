package application

import (
	"github.com/google/uuid"

	"github.com/ardihikaru/go-chi-example-part-1/pkg/config"
	"github.com/ardihikaru/go-chi-example-part-1/pkg/jwtauth"
	"github.com/ardihikaru/go-chi-example-part-1/pkg/logger"
)

// Dependencies hold the primitives and structs and/or interfaces that are required
// for the application's business logic.
type Dependencies struct {
	SvcId     string
	Cfg       *config.Config
	Log       *logger.Logger
	TokenAuth *jwtauth.JWTAuth
}

// BuildDependencies builds dependencies
func BuildDependencies(cfg *config.Config, log *logger.Logger) *Dependencies {
	// generates service ID
	svcId := uuid.New().String()

	// initializes JWT Authenticator
	tokenAuth := getTokenAuthentication(&cfg.JwtAuth, log)

	return &Dependencies{
		SvcId:     svcId,
		Cfg:       cfg,
		Log:       log,
		TokenAuth: tokenAuth,
	}
}
