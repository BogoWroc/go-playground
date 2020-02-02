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
	tp = template.Must(template.ParseGlob(dir + "/sandpit/04_web/level2_templates/nestingTemplates/templates/*.gohtml"))
}

func main() {
	err := tp.ExecuteTemplate(os.Stdout, "hello.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
