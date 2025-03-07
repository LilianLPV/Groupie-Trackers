package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Pokemonrecherche struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Supertypes string   `json:"supertype"`
	Subtypes   []string `json:"subtypes"`
	Types      []string `json:"types"`
	Images     struct {
		Small string `json:"small"`
	} `json:"images"`
}

type Requete struct {
	Data       []Pokemonrecherche `json:"data"`
	Page       int                `json:"page"`
	PageSize   int                `json:"pageSize"`
	Count      int                `json:"count"`
	TotalCount int                `json:"totalCount"`
}

func Recherche(query string, page int, pageSize int) (Requete, int, error) {
	// URL avec les paramètres de pagination
	urlApi := fmt.Sprintf("https://api.pokemontcg.io/v2/cards?q=name:%s*&select=id,name,images,supertype,subtypes,types&page=%d&pageSize=%d", query, page, pageSize)

	req, err := http.NewRequest("GET", urlApi, nil)
	if err != nil {
		return Requete{}, http.StatusInternalServerError, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("X-Api-Key", "38bfad3b-efb1-44bd-9fb3-a2ed677fc716")

	httpClient := http.Client{Timeout: 15 * time.Second}

	res, err := httpClient.Do(req)
	if err != nil {
		return Requete{}, http.StatusInternalServerError, fmt.Errorf("error executing request: %v", err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return Requete{}, res.StatusCode, fmt.Errorf("Erreur votre demande n'est pas bien formatée : %s", query)
	}

	var pokemonrequete Requete
	err = json.NewDecoder(res.Body).Decode(&pokemonrequete)
	if err != nil {
		return Requete{}, http.StatusInternalServerError, fmt.Errorf("error decoding JSON: %v", err)
	}

	return pokemonrequete, res.StatusCode, nil
}
