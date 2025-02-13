package handlers

import (
	"net/http"
	"sync"
	"text/template"
)

var tmplFile_correct = "template/answer.html.tmpl"

func HandleAnswer() http.Handler {
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

			answer := r.FormValue("answer")
			if answer != "World wide web" {
				err := tpl.ExecuteTemplate(w, "wrong", nil)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				return
			}

			err := tpl.ExecuteTemplate(w, "correct", nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		},
	)
}
