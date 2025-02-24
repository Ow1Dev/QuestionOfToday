package handlers

import (
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/Ow1Dev/QustionOfToday/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

var tmplFile = "template/question.html.tmpl"

var (
	question = "In a website browser address bar, what does “www” stand for?"
	id       = "00000000-0000-0000-0000-000000000000"
)

type Question struct {
	Question string
	Id       string
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
				tpl, tplerr = template.ParseFiles("template/base.html.tmpl", tmplFile)
			})

			var qu = Question{
				question,
				id,
			}
			if !debug {
				currettime := pgtype.Date{Time: time.Now(), Valid: true}
				question, err := repo.GetTodaysQuestion(r.Context(), currettime)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				qu = Question{
					question.Question,
					question.ID.String(),
				}
			}

			if tplerr != nil {
				http.Error(w, tplerr.Error(), http.StatusInternalServerError)
				return
			}

			err := tpl.ExecuteTemplate(w, "base", qu)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		},
	)
}
