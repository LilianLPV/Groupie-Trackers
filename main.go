package main

import (
	"fmt"
	"groupie-tracker/utils"
	"net/http"
	"slices"
	"strconv"
)

func main() {
	utils.InitTemplate()
	http.HandleFunc("/", AccueilHandler)
	http.HandleFunc("/recherche", RechercheHandler)
	http.HandleFunc("/favorites", FavorisHandler)
	http.HandleFunc("/erreur", ErreurHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))
	fmt.Println("Le serveur est lancé http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func ErreurHandler(w http.ResponseWriter, r *http.Request) {
	statutCode := r.FormValue("statutCode")
	erreur := r.FormValue("erreur")

	type Erreur struct {
		StatutCode string
		Message    string
	}

	err := utils.Tpl.ExecuteTemplate(w, "erreur", Erreur{StatutCode: statutCode, Message: erreur})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	accueil, err := utils.Accueil()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching accueil: %v", err), http.StatusInternalServerError)
		return
	}

	err = utils.Tpl.ExecuteTemplate(w, "accueil", accueil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func RechercheHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("search")
	types := r.FormValue("type")
	supertypes := r.FormValue("supertypes")
	subtypes := r.FormValue("subtypes")

	// Récupération des paramètres de pagination avec des valeurs par défaut
	page := 1
	pageSize := 20 // ou tout autre taille de page par défaut

	// Si une page spécifique est demandée dans l'URL
	if pageStr := r.FormValue("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	// Si une taille de page spécifique est demandée
	if pageSizeStr := r.FormValue("pageSize"); pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
			pageSize = ps
		}
	}

	fmt.Println(query, types, supertypes, subtypes, page, pageSize)

	listResult, codeResult, err := utils.Recherche(query, page, pageSize)
	if err != nil || codeResult != http.StatusOK {
		http.Redirect(w, r, fmt.Sprintf("/erreur?statutCode=%d&erreur=%s", codeResult, err.Error()), http.StatusSeeOther)
		return
	}

	// Filtrage des résultats
	var listResultFiltered utils.Requete
	if supertypes == "supertype" && subtypes == "subtype" && types == "type" {
		listResultFiltered = listResult
	} else {
		// Initialisation de la structure filtrée
		listResultFiltered = utils.Requete{
			Page:       listResult.Page,
			PageSize:   listResult.PageSize,
			Count:      0, // Sera mis à jour à la fin du filtrage
			TotalCount: listResult.TotalCount,
		}

		for _, pokemon := range listResult.Data {
			var checkTypes bool = (types == "type" || slices.Contains(pokemon.Types, types))
			var checkSubtypes bool = (subtypes == "subtype" || slices.Contains(pokemon.Subtypes, subtypes))
			var checkSupertypes bool = (supertypes == "supertype" || pokemon.Supertypes == supertypes)

			if checkTypes && checkSupertypes && checkSubtypes {
				listResultFiltered.Data = append(listResultFiltered.Data, pokemon)
			}
		}

		// Mettre à jour le nombre de résultats filtrés
		listResultFiltered.Count = len(listResultFiltered.Data)
	}

	// Préparer les données de pagination pour le template
	data := struct {
		Results     utils.Requete
		CurrentPage int
		TotalPages  int
		HasPrevPage bool
		HasNextPage bool
		Query       string
		Types       string
		Supertypes  string
		Subtypes    string
	}{
		Results:     listResultFiltered,
		CurrentPage: page,
		TotalPages:  (listResult.TotalCount + pageSize - 1) / pageSize, // Calcule le nombre total de pages
		HasPrevPage: page > 1,
		HasNextPage: page*pageSize < listResult.TotalCount,
		Query:       query,
		Types:       types,
		Supertypes:  supertypes,
		Subtypes:    subtypes,
	}

	err = utils.Tpl.ExecuteTemplate(w, "recherche.html", data)
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

	err := utils.Tpl.ExecuteTemplate(w, "filtre.html", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}
