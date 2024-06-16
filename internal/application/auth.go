package application

import (
	"github.com/ardihikaru/go-chi-example-part-1/pkg/authenticator"
	"github.com/ardihikaru/go-chi-example-part-1/pkg/config"
	"github.com/ardihikaru/go-chi-example-part-1/pkg/jwtauth"
	"github.com/ardihikaru/go-chi-example-part-1/pkg/logger"
	e "github.com/ardihikaru/go-chi-example-part-1/pkg/utils/error"
)

// getTokenAuthentication creates an authentication token from the authenticator
func getTokenAuthentication(jwtCfg *config.JwtAuth, log *logger.Logger) *jwtauth.JWTAuth {
	tokenAuth, err := authenticator.MakeTokenAuth(jwtCfg.JWTAlgorithm, jwtCfg.JWTSecret)
	if err != nil {
		e.FatalOnError(err, "failed to create a JWT authenticator")
	}

	return tokenAuth
}
