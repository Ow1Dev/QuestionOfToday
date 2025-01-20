package server

import (
	"net/http"
	"sync"
	"text/template"

	"github.com/Ow1Dev/QustionOfToday/internal/config"
	"github.com/Ow1Dev/QustionOfToday/internal/middleware"
)

type Question struct {
	Question string
}

var tmplFile = "template/question.html.tmpl"

var q = Question{
	Question: "In a website browser address bar, what does “www” stand for?",
}

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
  mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	mux.Handle("/", handleIndex())
}

func handleIndex() http.Handler {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
      init.Do(func(){
        tpl, tplerr = template.ParseFiles(tmplFile)
      })

			if tplerr != nil {
				http.Error(w, tplerr.Error(), http.StatusInternalServerError)
				return
			}

      err := tpl.ExecuteTemplate(w, "question", q)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
        return
			}
		},
	)
}
