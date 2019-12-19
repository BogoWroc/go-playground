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
	tp = template.Must(template.ParseFiles(dir + "/sandpit/04_web/level2_templates/passingCompositeDataIntoTemplate/hello.gohtml"))
}

type person struct {
	Name string
	Age  int
}

func main() {
	bob := person{
		Name: "Bob",
		Age:  23,
	}
	alice := person{
		Name: "Alice",
		Age:  32,
	}
	d := []person{bob, alice}

	err := tp.ExecuteTemplate(os.Stdout, "hello.gohtml", d)
	if err != nil {
		log.Fatalln(err)
	}
}
