package handlers

import (
	"fmt"
	"net/http"
	"sync"
	"text/template"

	"github.com/Ow1Dev/QustionOfToday/internal/repository"
)

var tmplFile = "template/question.html.tmpl"

type question struct {
	Question string
}

var q = question{
	Question: "In a website browser address bar, what does “www” stand for?",
}

func HandleIndex(
	repo *repository.Queries,
) http.Handler {
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

      question, err := repo.GetAllQuestions(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

      fmt.Println(question[0])

			if tplerr != nil {
				http.Error(w, tplerr.Error(), http.StatusInternalServerError)
				return
			}

      err = tpl.ExecuteTemplate(w, "question", question[0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
        return
			}

		},
	)
}
