// Package logger provides functions to set up a new logger
package logger

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	mw "github.com/ardihikaru/go-chi-example-part-1/pkg/middleware"
)

var prefix string
var reqId uint64

const (
	logFormatText    = "text"
	logFormatConsole = "console"
)

// Logger is a small wrapper around a zap.Logger.
type Logger struct {
	*zap.Logger
}

// init initializes prefix value
// FYI: this function adopts go-chi middleware
func init() {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		_, err := rand.Read(buf[:])
		if err != nil {
			return
		}
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}

	prefix = fmt.Sprintf("%s/%s", hostname, b64[0:10])
}

// New creates a new Logger with given logLevel and logFormat as part of a permanent field of the logger.
func New(logLevel, logFormat string) (*Logger, error) {
	if logFormat == logFormatText {
		logFormat = logFormatConsole
	}

	zapConfig := zap.NewProductionConfig()
	zapConfig.Encoding = logFormat

	var level zapcore.Level
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		return nil, err
	}
	zapConfig.Level = zap.NewAtomicLevelAt(level)
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("could not build logger: %w", err)
	}

	zap.ReplaceGlobals(logger)

	return &Logger{Logger: logger}, nil
}

// getRequestId is a middleware that injects a request ID into the context of each
// request. A request ID is a string of the form "host.example.com/random-0001",
// where "random" is a base62 random string that uniquely identifies this go
// process, and where the last number is an atomically incremented request
// counter.
// FYI: this function adopts go-chi middleware
func getRequestId(r *http.Request) string {
	requestId := r.Header.Get(mw.RequestId)
	if requestId == "" {
		myId := atomic.AddUint64(&reqId, 1)
		requestId = fmt.Sprintf("%s-%06d", prefix, myId)
	}

	return requestId
}

// SetLogger returns a middleware that logs the start and end of each request, along
// with some useful data about what was requested, what the response status was,
// and how long it took to return.
// Inspired by https://github.com/treastech/logger
func SetLogger(l *Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				l.Info("served request",
					zap.String("proto", r.Proto),
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.Duration("lat", time.Since(t1)),
					zap.Int("status", ww.Status()),
					zap.Int("size", ww.BytesWritten()),
					zap.String("latencyStr", (time.Since(t1)).String()),
					zap.Duration("latencyDuration", time.Since(t1)),
					zap.String("reqId", getRequestId(r)))
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
