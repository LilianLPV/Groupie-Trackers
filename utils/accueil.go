package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Pokemon struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Supertype   string   `json:"supertype"`
	Subtype     []string `json:"subtype"`
	Level       string   `json:"level"`
	HP          string   `json:"hp"`
	Types       []string `json:"types"`
	EvolvesFrom string   `json:"evolvesFrom"`
	Abilities   []struct {
		Name string `json:"name"`
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"abilities"`
	Attacks []struct {
		Name                string   `json:"name"`
		Cost                []string `json:"cost"`
		ConvertedEnergyCost int      `json:"convertedEnergyCost"`
		Damage              string   `json:"damage"`
		Text                string   `json:"text"`
	} `json:"attacks"`
	Weaknesses []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"weaknesses"`
	Resistances []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"resistances"`
	RetreatCost          []string `json:"retreatCost"`
	ConvertedRetreatCost int      `json:"convertedRetreatCost"`
	Set                  struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Series       string `json:"series"`
		PrintedTotal int    `json:"printedTotal"`
		Total        int    `json:"total"`
		Legalities   struct {
			Unlimited string `json:"unlimited"`
		} `json:"legalities"`
		PtcgoCode   string `json:"ptcgoCode"`
		ReleaseDate string `json:"releaseDate"`
		UpdatedAt   string `json:"updatedAt"`
		Images      struct {
			Symbol string `json:"symbol"`
			Logo   string `json:"logo"`
		} `json:"images"`
	} `json:"set"`
	Number                 string `json:"number"`
	Artist                 string `json:"artist"`
	Rarity                 string `json:"rarity"`
	FlavorText             string `json:"flavorText"`
	NationalPokedexNumbers []int  `json:"nationalPokedexNumbers"`
	Legalities             struct {
		Unlimited string `json:"unlimited"`
	} `json:"legalities"`
	Images struct {
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"images"`
}

type PokemonResponse struct {
	Data []Pokemon `json:"data"`
}

func Accueil() ([]Pokemon, error) {
	urlApi := "https://api.pokemontcg.io/v2/cards/"

	req, err := http.NewRequest("GET", urlApi, nil)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("X-Api-Key", "38bfad3b-efb1-44bd-9fb3-a2ed677fc716")

	httpClient := http.Client{Timeout: 15 * time.Second}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var pokemonResponse PokemonResponse
	err = json.NewDecoder(res.Body).Decode(&pokemonResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return pokemonResponse.Data, nil
}
