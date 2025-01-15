package server

import (
	"net/http"

	"github.com/Ow1Dev/QustionOfToday/internal/config"
	"github.com/Ow1Dev/QustionOfToday/internal/middleware"
)

func NewServer(
	config *config.Config,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux)

	var handler http.Handler = mux
	handler = logging.NewLoggingMiddleware(handler)
	return handler
}

func addRoutes(
	mux *http.ServeMux,
) {
	mux.Handle("/", http.NotFoundHandler())
}
