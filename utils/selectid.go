package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CardItem struct {
	Data struct {
		ID     string   `json:"id"`
		Name   string   `json:"name"`
		Types  []string `json:"types"`
		Images struct {
			Small string `json:"small"`
		} `json:"images"`
	} `json:"data"`
}

func Selectid(id string) (CardItem, error) {
	urlApi := "https://api.pokemontcg.io/v2/cards/" + id

	req, err := http.NewRequest("GET", urlApi, nil)

	if err != nil {
		return CardItem{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("X-Api-Key", "38bfad3b-efb1-44bd-9fb3-a2ed677fc716")

	httpClient := http.Client{Timeout: 15 * time.Second}

	res, err := httpClient.Do(req)
	if err != nil {
		return CardItem{}, fmt.Errorf("error executing request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return CardItem{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var pokemonResponse CardItem
	err = json.NewDecoder(res.Body).Decode(&pokemonResponse)
	if err != nil {
		return CardItem{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	return pokemonResponse, nil
}
