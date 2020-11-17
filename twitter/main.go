package main

import (
	"log"
	"os"
	"text/template"
)

type midia struct {
	Avatar string
	Capa   string
}

type tweetConfig struct {
	Tweets      float64
	Notificação bool
	Follow      bool
	DataDaConta string
	Seguindo    float64
	Seguidores  float64
}

type perfil struct {
	Nome        string
	User        string
	Bio         string
	Localização string
	Site        string
	midia
	tweetConfig
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("perfil/*"))
}

func main() {
	tecnoGueto := perfil{
		"Tecnogueto",
		"tecnogueto",
		`Nós da Tecnogueto, queremos socializar o cenário atual, ensinando profissões 
dentro da tecnologia e promovendo a diversidade de gênero, raça/etnia e cultural.`,
		"Rio de Janeiro, Brasil",
		"tecnogueto.com.br",
		midia{"https://pbs.twimg.com/profile_images/1162103759218139137/CYfSOQFx_400x400.jpg", "https://pbs.twimg.com/profile_banners/1002022507413757958/1553770868/1500x500"},
		tweetConfig{
			448,
			true,
			true,
			"Maio de 2018",
			100,
			3.2343,
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "perfil.txt", tecnoGueto)
	if err != nil {
		log.Fatalln(err)
	}
}
