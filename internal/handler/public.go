package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"github.com/ardihikaru/go-chi-example-part-2/pkg/logger"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/utils/http"
)

// PublicHandler handle public routes
func PublicHandler(serviceId string, log *logger.Logger) http.Handler {
	//func PublicHandler(serviceId string, log *logger.Logger, timeout time.Duration) http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.HandleFunc("/service-id", getServiceId(serviceId)) // GET /roles - Read a list of users.

		r.Route("/with-sleep", func(r chi.Router) {
			r.HandleFunc("/", getResponseWithSleep(log)) // GET /roles - Read a list of users.
		})

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

// getServiceId gets serviceId
// @Summary 	This API can be used as health check for this application
// @Description Tells the service ID of this service.
// @Tags 		public
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} httputil.Response "api response"
// @Router 		/public/with-sleep [get]
func getResponseWithSleep(log *logger.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// adds sleep for 2 seconds
		// here, we expect that the api service will ONLY respond after 2 seconds of delay,
		// e.g.
		//	if we set the config as http.timeout=4s, then, it will respond properly
		//	if we set the config as http.timeout=1s, then, it will respond a timeout (503)
		//time.Sleep(6 * time.Second)

		//// first alt.: using SLEEP
		log.Info("do something here .... sleeping in 2 seconds ...")
		time.Sleep(2 * time.Second)

		// second alt.: using CONTEXT
		ctx := r.Context()
		////processTime := 6 * time.Second
		//processTime := 3 * time.Second

		select {

		// got a timeout
		case <-ctx.Done():
			log.Info("masuk case DONE ...")

			return

		default:
			log.Info("done ...")
			//case <-time.After(processTime):
			//	log.Info("after 2 seconds delay ... return 200 (OK)")

			//return
		}

		log.Info("after SELECT-CASE")

		// The above channel simulates some hard work.
		_ = httputil.WriteResponse(w, httputil.SuccessResponse, &httputil.Response{
			StatusCode: 200,
		})

		return
	}
}
