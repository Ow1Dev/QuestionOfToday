package handlers

import (
	"net/http"
	"sync"
	"text/template"
)

var tmplFile_correct = "template/anwser.html.tmpl"

func HandleAnwser() http.Handler {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
      init.Do(func(){
        tpl, tplerr = template.ParseFiles(tmplFile_correct)
      })

			if tplerr != nil {
				http.Error(w, tplerr.Error(), http.StatusInternalServerError)
				return
			}

      anwser := r.FormValue("anwser")
      if anwser != "World wide web" {
        err := tpl.ExecuteTemplate(w, "worng", q)
        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }

        return
      }

      err := tpl.ExecuteTemplate(w, "correct", q)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
        return
			}
		},
	)
}
