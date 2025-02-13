package server

import (
	"net/http"

	"github.com/Ow1Dev/QustionOfToday/internal/config"
	"github.com/Ow1Dev/QustionOfToday/internal/handlers"
	"github.com/Ow1Dev/QustionOfToday/internal/middleware"
	"github.com/Ow1Dev/QustionOfToday/internal/repository"
	"github.com/rs/zerolog"
)

func NewServer(
	logger *zerolog.Logger,
	config *config.Config,
	repo *repository.Queries,
	debug bool,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, repo, debug)
	var handler http.Handler = mux
	handler = logging.NewLoggingMiddleware(logger, handler)
	return handler
}

func addRoutes(
	mux *http.ServeMux,
	repo *repository.Queries,
	debug bool,
) {
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	mux.Handle("POST /answer", handlers.HandleAnswer())
	mux.Handle("GET /", handlers.HandleIndex(repo, debug))
}
