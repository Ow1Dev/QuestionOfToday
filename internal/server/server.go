package server

import (
	"net/http"

	"github.com/Ow1Dev/QustionOfToday/internal/config"
	"github.com/Ow1Dev/QustionOfToday/internal/handlers"
	"github.com/Ow1Dev/QustionOfToday/internal/middleware"
	"github.com/rs/zerolog"
)

func NewServer(
  logger *zerolog.Logger,
	config *config.Config,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux)

	var handler http.Handler = mux
	handler = logging.NewLoggingMiddleware(logger, handler)
	return handler
}

func addRoutes(
	mux *http.ServeMux,
) {
  mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	mux.Handle("/anwser", handlers.HandleAnwser())
	mux.Handle("/", handlers.HandleIndex())
}

