package handlers

import (
	"net/http"
	"sync"
	"text/template"
)

var tmplFile = "template/question.html.tmpl"

type question struct {
	Question string
}

var q = question{
	Question: "In a website browser address bar, what does “www” stand for?",
}

func HandleIndex() http.Handler {
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
