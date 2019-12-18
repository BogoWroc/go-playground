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
	tp = template.Must(template.ParseFiles(dir + "/sandpit/04_web/level2_templates/passingDataIntoTemplate/hello.gohtml"))
}

func main() {
	err := tp.ExecuteTemplate(os.Stdout, "hello.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
