package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	Cars := []car{f, c}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", Cars)
	if err != nil {
		log.Fatalln(err)
	}
}
