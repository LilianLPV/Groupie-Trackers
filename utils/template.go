package utils

import (
	"log"
	"text/template"
)

var Tpl *template.Template

func InitTemplate() {
	var err error
	Tpl, err = template.ParseGlob("template/*.html")
	if err != nil {
		log.Fatal("Erreur lors du chargement des templates :", err)
	}
}
