package main

import (
	"log"
	"os"
	"text/template"
)

var tp *template.Template

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tp = template.Must(template.ParseGlob(dir + "/sandpit/04_web/level2_templates/parsingAndExecuting/templates/*.gohtml"))
}

func main() {
	err := tp.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tp.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tp.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
