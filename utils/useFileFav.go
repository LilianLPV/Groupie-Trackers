package utils

import (
	"encoding/json"
	"os"
)

// Lecture du fichier
func ReadFileFav() ([]string, error) {
	fileData, fileErr := os.ReadFile("./data/fav.json")
	if fileErr != nil {
		return nil, fileErr
	}

	var listFavorits []string
	decodeErr := json.Unmarshal(fileData, &listFavorits)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return listFavorits, nil
}

// Ecriture du fichier
func WriteFileFav(listFavorits []string) error {
	fileDataJson, jsonErr := json.Marshal(listFavorits)
	if jsonErr != nil {
		return jsonErr
	}

	fileErr := os.WriteFile("./data/fav.json", fileDataJson, 0644)
	if fileErr != nil {
		return fileErr
	}

	return nil
}

// Supprssion d'une valeur



// Ajout d'une valeur
func AddFav(newItem string) error {
	fileData, fileErr := ReadFileFav()
	if fileErr != nil {
		return fileErr
	}

	// Il faut gerer les doublons...
	// voir doc package slice !
	fileData = append(fileData, newItem)

	fileErr = WriteFileFav(fileData)
	if fileErr != nil {
		return fileErr
	}

	return nil
}
