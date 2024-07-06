package handler

import (
	"fmt"
	"net/http"

	"github.com/developersismedika/sqlx"
	"github.com/go-chi/chi"

	eRs "github.com/ardihikaru/go-chi-example-part-2/internal/enum/resource"
	"github.com/ardihikaru/go-chi-example-part-2/internal/storage/resourcerolemap"

	"github.com/ardihikaru/go-chi-example-part-2/pkg/enforcer"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/jwtauth"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/logger"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/middleware"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/mysqldb"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/service/middlewareutility"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/service/session"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/utils/http"
)

type privateHandler struct {
	log *logger.Logger
}

// PrivateHandler handle private routes
func PrivateHandler(serviceId string, log *logger.Logger, tokenAuth *jwtauth.JWTAuth,
	enforcerPolicy *enforcer.Enforcer, db *sqlx.DB) http.Handler {
	r := chi.NewRouter()

	// builds resource group storage
	rsRoleStorage := &resourcerolemap.Storage{
		Storage: &mysqldb.Storage{Db: db, Log: log},
	}

	// initializes session middleware resource
	mwUtilSvc := middlewareutility.NewService(log, enforcerPolicy, rsRoleStorage)
	sessionSvc := session.NewService(log)
	mw := middleware.NewMiddleware(mwUtilSvc, sessionSvc)

	controller := privateHandler{log: log}

	r.Route("/", func(r chi.Router) {
		// Seeks, verifies and validates JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// validates token. Got invalids if (expired, missing)
		r.Use(jwtauth.Authenticator)

		// extracts the session on the URL parameter
		r.Use(mw.SessionCtx)

		// authorizes access control
		r.Use(mw.AuthorizeAccess(eRs.User, "read"))
		r.HandleFunc("/service-id", controller.getServiceId(serviceId))
	})

	return r
}

// getServiceId gets serviceId privately
// @Summary 	This API can be used as health check for this application
// @Description Tells the service ID of this service.
// @Tags 		public
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} httputil.Response "api response"
// @Router 		/private/service-id [get]
func (h *privateHandler) getServiceId(serviceId string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionData := r.Context().Value(middleware.SessionKey).(session.Session)
		h.log.Debug(fmt.Sprintf("session username: %s", sessionData.Username))

		_ = httputil.WriteResponse(w, httputil.SuccessResponse, &httputil.Response{
			Data:       serviceId,
			StatusCode: 200,
		})
		return
	}
}
