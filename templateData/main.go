package main

import (
	"log"
	"os"
	"text/template"
)

type data struct {
	Nome  string
	Idade int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("data/*"))
}

func main() {

	perfil := data{
		"vapordev",
		29,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "dados_pessoais.txt", perfil)
	if err != nil {
		log.Fatalln(err)
	}
}
