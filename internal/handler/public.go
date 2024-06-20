package handler

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ardihikaru/go-chi-example-part-1/pkg/utils/http"
)

// PublicHandler handle public routes
func PublicHandler(serviceId string) http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.HandleFunc("/service-id", getServiceId(serviceId)) // GET /roles - Read a list of users.
	})

	return r
}

// getServiceId gets serviceId
// @Summary 	This API can be used as health check for this application
// @Description Tells the service ID of this service.
// @Tags 		public
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} httputil.Response "api response"
// @Router 		/public/service-id [get]
func getServiceId(serviceId string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = httputil.WriteResponse(w, httputil.SuccessResponse, &httputil.Response{
			Data:       serviceId,
			StatusCode: 200,
		})
		return
	}
}
