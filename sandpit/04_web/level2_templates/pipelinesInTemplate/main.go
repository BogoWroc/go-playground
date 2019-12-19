package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tp *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tp = template.Must(
		template.New("").
			Funcs(fm).
			ParseFiles(dir + "/sandpit/04_web/level2_templates/pipelinesInTemplate/hello.gohtml"))
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
