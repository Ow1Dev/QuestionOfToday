package handlers

import (
	"net/http"
	"strings"
	"sync"
	"text/template"

	"github.com/Ow1Dev/QustionOfToday/internal/repository"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

var tmplFile_correct = "template/answer.html.tmpl"

func HandleAnswer(repo *repository.Queries, logger *zerolog.Logger) http.Handler {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			init.Do(func() {
				tpl, tplerr = template.ParseFiles(tmplFile_correct)
			})

			if tplerr != nil {
				http.Error(w, tplerr.Error(), http.StatusInternalServerError)
				return
			}

			userAnswer := r.FormValue("answer")
			id := r.FormValue("id")

			logger.Debug().Msgf("Retrieving answer with id %s", id)

			if id == "" {
				http.NotFound(w, r)
				return
			}

			uuidid, err := uuid.Parse(id)
			if err != nil {
				http.Error(w, tplerr.Error(), http.StatusInternalServerError)
				return
			}

			realanswer, err := repo.GetAnswerById(r.Context(), uuidid)
			if err != nil {
				http.Error(w, tplerr.Error(), http.StatusInternalServerError)
				return
			}

			if strings.ToLower(userAnswer) != realanswer.Answer {
				err := tpl.ExecuteTemplate(w, "wrong", nil)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				return
			}

			err = tpl.ExecuteTemplate(w, "correct", realanswer)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		},
	)
}
