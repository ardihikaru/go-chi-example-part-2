package timeouthandler

import "github.com/ardihikaru/go-chi-example-part-1/pkg/logger"

type Service struct {
	log *logger.Logger
}

func (s Service) LogInfo(msg string) {
	s.log.Info(msg)
}

func (s Service) LogWarn(msg string) {
	s.log.Warn(msg)
}

func (s Service) LogError(msg string) {
	s.log.Error(msg)
}

func (s Service) LogDebug(msg string) {
	s.log.Debug(msg)
}

// NewService creates a new timeout handler service
func NewService(log *logger.Logger) *Service {
	return &Service{
		log: log,
	}
}
