package handlers

import (
	"net/http"
	"sync"
	"text/template"

	"github.com/Ow1Dev/QustionOfToday/internal/repository"
)

var tmplFile = "template/question.html.tmpl"

var (
	question = "In a website browser address bar, what does “www” stand for?"
)

type Question struct {
	Question string
}

func HandleIndex(
	repo *repository.Queries,
	debug bool,
) http.Handler {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			init.Do(func() {
				tpl, tplerr = template.ParseFiles(tmplFile)
			})

			var qu = question

			if !debug {
				question, err := repo.GetAllQuestions(r.Context())
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				qu = question[0].Question
			}

			if tplerr != nil {
				http.Error(w, tplerr.Error(), http.StatusInternalServerError)
				return
			}

			err := tpl.ExecuteTemplate(w, "question", Question{
				Question: qu,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		},
	)
}
