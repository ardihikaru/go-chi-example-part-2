package handler

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ardihikaru/go-chi-example-part-1/internal/service/auth"

	"github.com/ardihikaru/go-chi-example-part-1/pkg/config"
	"github.com/ardihikaru/go-chi-example-part-1/pkg/jwtauth"
	"github.com/ardihikaru/go-chi-example-part-1/pkg/logger"
	"github.com/ardihikaru/go-chi-example-part-1/pkg/utils/http"
)

// handler defines the struct to wrap handlers
type handler struct {
	log     *logger.Logger
	authSvc *auth.Service
}

type authResponse struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status,omitempty"`
	Data       *auth.Token `json:"data,omitempty"`
}

// AuthHandler handles all authentication routes
func AuthHandler(cfg *config.Config, log *logger.Logger, tokenAuth *jwtauth.JWTAuth) http.Handler {
	r := chi.NewRouter()

	// initializes services
	authSvc := auth.NewService(log, cfg.JwtAuth.JWTExpiredInSec, tokenAuth)

	h := handler{
		log:     log,
		authSvc: authSvc,
	}

	r.Route("/login", func(r chi.Router) {
		r.Post("/", h.authLogin()) // POST /auth/login - authorize login user
	})

	return r
}

// authLogin processes the request to create access token
func (h *handler) authLogin() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var errMsg *string
		var payload auth.LoginData

		// extracts request body
		errMsg, httpCode, err := httputil.GetJsonBody(r.Body, &payload)
		if err != nil {
			_ = httputil.WriteResponse(w, httputil.ErrorResponse, &httputil.ErrResponse{
				StatusCode: httpCode, Error: err, Message: *errMsg,
			})
			return
		}

		// authorizes request
		tokenCred, err := h.authSvc.Authorize(payload)
		if err != nil {
			_ = httputil.WriteResponse(w, httputil.ErrorResponse, &httputil.ErrResponse{
				StatusCode: http.StatusUnauthorized, Error: err, Message: "unauthorized access",
			})
			return
		}

		resp := authResponse{
			StatusCode: 200,
			Data:       tokenCred,
		}

		_ = httputil.WriteResponse(w, httputil.SuccessResponse, &resp)
	}
}
