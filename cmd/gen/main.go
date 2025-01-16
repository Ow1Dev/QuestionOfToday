package main

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

type Question struct {
	Question string
}

var tmplFile = "template/question.html.tmpl"

var q = Question{
	Question: "In a website browser address bar, what does “www” stand for?",
}

func main() {
	if err := run(os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(stdout io.Writer, args []string) error {
	_ = stdout
	_ = args

	if _, err := os.Stat(tmplFile); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", tmplFile)
	}

	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		return err
	}

	err = writeToFile(tmpl, q)
	if err != nil {
		return err
	}

	return nil
}

func writeToFile(tmpl *template.Template, q Question) error {
	var f *os.File
	f, err := os.Create("static/index.html")
	if err != nil {
		return err
	}
	err = tmpl.ExecuteTemplate(f, "question", q)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
