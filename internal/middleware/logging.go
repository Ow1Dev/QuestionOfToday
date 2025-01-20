package logging

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

func NewLoggingMiddleware(logger *zerolog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

		next.ServeHTTP(w, r)

    logger.Info().
      Str("method", r.Method).
      Str("url", r.URL.RequestURI()).
      Dur("duration", time.Since(start)).
      Msg("incoming request")
	})
}
