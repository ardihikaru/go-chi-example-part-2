// Package router provides the base configurations to build a router.
// This package is open for an extension whether new routes need to be added.
package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/ardihikaru/go-chi-example-part-2/internal/application"
	"github.com/ardihikaru/go-chi-example-part-2/internal/handler"
	"github.com/ardihikaru/go-chi-example-part-2/internal/service/session"
	"github.com/ardihikaru/go-chi-example-part-2/internal/storage/resourcerolemap"

	"github.com/ardihikaru/go-chi-example-part-2/pkg/logger"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/middleware"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/mysqldb"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/service/middlewareutility"
)

// GetRouter configures a chi router and starts the http server
func GetRouter(deps *application.Dependencies) *chi.Mux {
	r := chi.NewRouter()

	if deps.Log != nil {
		r.Use(logger.SetLogger(deps.Log))
	}

	// builds resource group storage
	rsRoleStorage := &resourcerolemap.Storage{
		Storage: &mysqldb.Storage{Db: deps.Db, Log: deps.Log},
	}

	// builds middleware
	mwUtilSvc := middlewareutility.NewService(deps.Log, deps.Enforcer, rsRoleStorage)
	sessionSvc := session.NewService(deps.Log)
	mw := middleware.NewMiddleware(mwUtilSvc, sessionSvc)

	r.Use(mw.Timeout(deps.Cfg.Http.WriteTimeout)) // returns 504

	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   deps.Cfg.Cors.AllowedOrigins,
		AllowedMethods:   deps.Cfg.Cors.AllowedMethods,
		AllowedHeaders:   deps.Cfg.Cors.AllowedHeaders,
		ExposedHeaders:   deps.Cfg.Cors.ExposedHeaders,
		AllowCredentials: deps.Cfg.Cors.AllowCredentials,
		MaxAge:           deps.Cfg.Cors.MaxAge,
		Debug:            deps.Cfg.Cors.Debug,
	}))

	buildTree(r, deps)

	return r
}

func buildTree(r *chi.Mux, deps *application.Dependencies) {
	// handles swagger route
	r.Mount("/swagger", httpSwagger.WrapHandler)

	// handles public related route(s)
	r.Mount("/public", handler.PublicHandler(deps.SvcId, deps.Log))

	// handles private related route(s)
	r.Mount("/private", handler.PrivateHandler(deps.SvcId, deps.Log, deps.TokenAuth, deps.Enforcer, deps.Db))

	// handles auth related route(s)
	r.Mount("/auth", handler.AuthHandler(deps.Cfg, deps.Log, deps.TokenAuth))
}
