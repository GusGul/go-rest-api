package db

import (
	"encoding/json"
	"go-rest-api/pkg/models"
	"io/ioutil"
	"log"
)

var Albums []models.Album

func LoadDatabase(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read database file: %v", err)
	}

	err = json.Unmarshal(data, &Albums)
	if err != nil {
		log.Fatalf("Failed to unmarshal database file: %v", err)
	}
}

func SaveDatabase(filePath string) {
	data, err := json.MarshalIndent(Albums, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal database: %v", err)
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write database file: %v", err)
	}
}
