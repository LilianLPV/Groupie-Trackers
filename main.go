package main

import (
	"fmt"
	"groupie-tracker/utils"
	"net/http"
)

func main() {
	utils.InitTemplate()
	utils.Accueil()
	utils.Recherche()
	utils.Favoris()
	utils.Filtre()
	http.HandleFunc("/", AccueilHandler)
	http.HandleFunc("/recherche", RechercheHandler)
	http.HandleFunc("/favorites", FavorisHandler)
	http.HandleFunc("/filtre", FiltreHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))
	fmt.Println("Le serveur est lancé http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	accueil, err := utils.Accueil()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching accueil: %v", err), http.StatusInternalServerError)
		return
	}

	err = utils.Tpl.ExecuteTemplate(w, "accueil.html", accueil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func RechercheHandler(w http.ResponseWriter, r *http.Request) {
	accueil, err := utils.Recherche()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching accueil: %v", err), http.StatusInternalServerError)
		return
	}

	err = utils.Tpl.ExecuteTemplate(w, "recherche.html", accueil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func FavorisHandler(w http.ResponseWriter, r *http.Request) {
	accueil, err := utils.Favoris()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching accueil: %v", err), http.StatusInternalServerError)
		return
	}

	err = utils.Tpl.ExecuteTemplate(w, "favoris.html", accueil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func FiltreHandler(w http.ResponseWriter, r *http.Request) {
	accueil, err := utils.Filtre()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching accueil: %v", err), http.StatusInternalServerError)
		return
	}

	err = utils.Tpl.ExecuteTemplate(w, "filtre.html", accueil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}
