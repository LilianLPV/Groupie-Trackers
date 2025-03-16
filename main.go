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
	http.HandleFunc("/search", RechercheHandler)
	http.HandleFunc("/favorites", FavorisHandler)
	http.HandleFunc("/favorites/add", FavoritesAdd)
	http.HandleFunc("/erreur", ErreurHandler)
	http.HandleFunc("/aboutit", AproposHandler)
	http.HandleFunc("/selectid", SelectIDHandler)
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
	// recupe liste id
	fileData, fileErr := utils.ReadFileFav()
	if fileErr != nil {
		http.Redirect(w, r, fmt.Sprintf("/erreur?statutCode=%d&erreur=%s", http.StatusInternalServerError, fileErr.Error()), http.StatusSeeOther)
		return
	}

var liste []utils.CardItem
	for _, id := range fileData { 
		card, err := utils.Selectid(id)
		if err != nil {
			http.Redirect(w, r, fmt.Sprintf("/erreur?statutCode=%d&erreur=%s", http.StatusInternalServerError, err.Error()), http.StatusSeeOther)
			return
		}
		liste = append(liste, card)
	}

	// bouclé su les la liste des id pour recup les cartes
	// for _, id := range fileData ???
	// ajout au tableau liste carte


	// envoie de la liste au template, changer nil par la variable ? 
	err := utils.Tpl.ExecuteTemplate(w, "favoris", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func AproposHandler(w http.ResponseWriter, r *http.Request) {
	apropos, err := utils.Apropos()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching accueil: %v", err), http.StatusInternalServerError)
		return
	}

	err = utils.Tpl.ExecuteTemplate(w, "apropos", apropos)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func SelectIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	fmt.Println(id)

	idresult, err := utils.Selectid(id)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/erreur?statutCode=%d&erreur=%s", http.StatusInternalServerError, err.Error()), http.StatusSeeOther)
		return
	}

	err = utils.Tpl.ExecuteTemplate(w, "selectid", idresult)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func FavoritesAdd(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	fmt.Println(id)
	res, resErr := utils.ReadFileFav()
	if resErr != nil {
		http.Redirect(w, r, fmt.Sprintf("/erreur?statutCode=%d&erreur=%s", http.StatusInternalServerError, resErr.Error()), http.StatusSeeOther)
		return
	}

	fmt.Println(res)

	errWrite := utils.AddFav(id)
	if errWrite != nil {
		http.Redirect(w, r, fmt.Sprintf("/erreur?statutCode=%d&erreur=%s", http.StatusInternalServerError, errWrite.Error()), http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}
