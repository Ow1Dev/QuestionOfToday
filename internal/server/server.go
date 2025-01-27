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
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, repo)
	var handler http.Handler = mux
	handler = logging.NewLoggingMiddleware(logger, handler)
	return handler
}

func addRoutes(
	mux *http.ServeMux,
	repo *repository.Queries,
) {
  mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	mux.Handle("/anwser", methodRestrict(http.MethodPost, handlers.HandleAnwser()))
	mux.Handle("/", methodRestrict(http.MethodGet, handlers.HandleIndex(repo)))
}

func methodRestrict(method string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != method {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
        next.ServeHTTP(w, r)
    })
}


