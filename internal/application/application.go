package application

import (
	"fmt"

	"github.com/developersismedika/sqlx"
	"github.com/google/uuid"

	"github.com/ardihikaru/go-chi-example-part-2/pkg/config"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/enforcer"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/jwtauth"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/logger"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/storage"
	e "github.com/ardihikaru/go-chi-example-part-2/pkg/utils/error"
)

// Dependencies hold the primitives and structs and/or interfaces that are required
// for the application's business logic.
type Dependencies struct {
	SvcId     string
	Cfg       *config.Config
	Db        *sqlx.DB
	Log       *logger.Logger
	TokenAuth *jwtauth.JWTAuth
	Enforcer  *enforcer.Enforcer
}

// BuildDependencies builds dependencies
func BuildDependencies(cfg *config.Config, log *logger.Logger) *Dependencies {
	// generates service ID
	svcId := uuid.New().String()

	// initializes JWT Authenticator
	tokenAuth := getTokenAuthentication(&cfg.JwtAuth, log)

	// initializes persistent store
	db, err := storage.DbConnect(log, cfg.DbMySQL)
	if err != nil {
		e.FatalOnError(err, "failed to open database connection")
	}

	// Load model configuration file and policy store adapter
	enforcerPolicy := enforcer.NewEnforcer(log, db, cfg.Enforcer)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	return &Dependencies{
		SvcId:     svcId,
		Cfg:       cfg,
		Db:        db,
		Log:       log,
		TokenAuth: tokenAuth,
		Enforcer:  enforcerPolicy,
	}
}
