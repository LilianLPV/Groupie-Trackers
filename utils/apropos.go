package utils

import (
	"fmt"
	"net/http"
	"time"
)

func Apropos() (string, error) {
	urlApi := "https://api.pokemontcg.io/v2/cards"

	req, err := http.NewRequest("GET", urlApi, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	httpClient := http.Client{Timeout: 15 * time.Second}
	res, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error executing request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	return "", nil
}
