package utils

import (
    "html/template"
    "log"
)

var Tpl *template.Template

func InitTemplate() {
    funcMap := template.FuncMap{
        "add": func(a, b int) int {
            return a + b
        },
        "subtract": func(a, b int) int {
            return a - b
        },
    }
    
    // Initialiser le template avec les fonctions personnalisées
    var err error
    Tpl = template.New("").Funcs(funcMap)
    
    // Charger vos templates - vérifiez que le chemin est correct
    Tpl, err = Tpl.ParseGlob("template/*.html") // Ajustez ce chemin selon votre structure de projet
    if err != nil {
        log.Fatalf("Erreur lors du chargement des templates: %v", err)
    }
}