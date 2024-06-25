package middlewareutility

import (
	"fmt"
	"net/http"

	"go.uber.org/zap/zapcore"

	"github.com/ardihikaru/go-chi-example-part-2/internal/service/session"

	"github.com/ardihikaru/go-chi-example-part-2/pkg/enforcer"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/logger"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/middleware"
)

// storage provides the interface for the functionality to fetch information from the table resource_group in the DB
type storage interface {
	GetObjListOwner(resource, role string) ([]string, error)
}

type Service struct {
	*enforcer.Enforcer
	log     *logger.Logger
	storage storage
}

// NewService creates a new timeout handler service
func NewService(log *logger.Logger, enforcerPolicy *enforcer.Enforcer, storage storage) *Service {
	return &Service{
		log:      log,
		Enforcer: enforcerPolicy,
		storage:  storage,
	}
}

// Log logs message based on the log level
func (svc *Service) Log(level zapcore.Level, msg string) {
	switch level {
	case zapcore.DebugLevel:
		svc.log.Debug(msg)
	case zapcore.InfoLevel:
		svc.log.Info(msg)
	case zapcore.WarnLevel:
		svc.log.Warn(msg)
	case zapcore.ErrorLevel:
		svc.log.Error(msg)
	case zapcore.FatalLevel:
		svc.log.Fatal(msg)
	case zapcore.PanicLevel:
		svc.log.Panic(msg)
	default:
		svc.log.Error(fmt.Sprintf("unexpected log level type: %d", level))
	}
}

// AuthorizeAccess is a middleware to enforce function's access control of the requester
func (svc *Service) AuthorizeAccess(resourceCode string, act string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var err error
			var ok bool
			//var objList []string

			// gets current user/subject
			sessionData := r.Context().Value(middleware.SessionKey).(session.Session)
			svc.log.Debug(fmt.Sprintf("captured subject (userId): %s", sessionData.Username))
			sub := sessionData.UserId

			//// loads object list from the resource code
			//objList, err = svc.storage.GetObjListOwner(resourceCode, sessionData.Role)
			//if err != nil {
			//	svc.log.Warn("error occurred when fetching object list")
			//	w.WriteHeader(http.StatusPreconditionFailed)
			//
			//	next.ServeHTTP(w, r)
			//	return
			//}
			//
			//// verifies access control in the resource level
			//if len(objList) == 0 {
			//	svc.log.Warn("this user has no access control into this resource")
			//	w.WriteHeader(http.StatusPreconditionFailed)
			//
			//	next.ServeHTTP(w, r)
			//	return
			//}

			// loads policy from Database
			err = svc.LoadPolicy()
			if err != nil {
				svc.log.Warn("failed to load enforcer policy")
				w.WriteHeader(http.StatusPreconditionFailed)

				next.ServeHTTP(w, r)
				return
			}

			// Casbin enforces policy for each object, whether the subject has an access or not
			ok, err = svc.Enforce(fmt.Sprint(sub), resourceCode, act)
			if err != nil {
				svc.log.Warn("error occurred when authorizing user")
				w.WriteHeader(http.StatusPreconditionFailed)

				next.ServeHTTP(w, r)
				return
			}
			if !ok {
				svc.log.Warn("this user is not authorized")
				w.WriteHeader(http.StatusUnauthorized)

				next.ServeHTTP(w, r)
				return
			}

			//// verifies access control in the action level
			//// Casbin enforces policy for each object, whether the subject has an access or not
			//for _, obj := range objList {
			//	ok, err = svc.Enforce(fmt.Sprint(sub), obj, act)
			//	if err != nil {
			//		svc.log.Warn("error occurred when authorizing user")
			//		w.WriteHeader(http.StatusPreconditionFailed)
			//
			//		next.ServeHTTP(w, r)
			//		return
			//	}
			//	if !ok {
			//		svc.log.Warn("this user is not authorized")
			//		w.WriteHeader(http.StatusUnauthorized)
			//
			//		next.ServeHTTP(w, r)
			//		return
			//	}
			//}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
