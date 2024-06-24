package middleware

import "net/http"

const (
	RequestId = "X-Request-Id"
)

// utility provides the interface for the functionality of logger.Logger and any other common utility
type utility interface {
	LogInfo(msg string)
	LogWarn(msg string)
	LogError(msg string)
	LogDebug(msg string)
	EnforcePolicy(sub, obj, act string) error
}

// session provides the interface for the functionality of session handler for the authentication
type session interface {
	SessionCtx(next http.Handler) http.Handler
}

type Resource struct {
	utility utility
	session session
}

// NewMiddleware creates a new middleware
func NewMiddleware(utility utility, session session) *Resource {
	return &Resource{
		utility: utility,
		session: session,
	}
}
