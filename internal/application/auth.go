package application

import (
	"github.com/ardihikaru/go-chi-example-part-2/pkg/authenticator"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/config"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/jwtauth"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/logger"
	e "github.com/ardihikaru/go-chi-example-part-2/pkg/utils/error"
)

// getTokenAuthentication creates an authentication token from the authenticator
func getTokenAuthentication(jwtCfg *config.JwtAuth, log *logger.Logger) *jwtauth.JWTAuth {
	tokenAuth, err := authenticator.MakeTokenAuth(jwtCfg.JWTAlgorithm, jwtCfg.JWTSecret)
	if err != nil {
		e.FatalOnError(err, "failed to create a JWT authenticator")
	}

	return tokenAuth
}
