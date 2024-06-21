package middleware

const (
	RequestId = "X-Request-Id"
)

// utility provides the interface for the functionality of logger.Logger
type utility interface {
	LogInfo(msg string)
	LogWarn(msg string)
	LogError(msg string)
	LogDebug(msg string)
}

type Resource struct {
	utility utility
}

// NewMiddleware creates a new middleware
func NewMiddleware(utility utility) *Resource {
	return &Resource{
		utility: utility,
	}
}
