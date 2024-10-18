package db

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "go-gin-api/pkg/models"
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