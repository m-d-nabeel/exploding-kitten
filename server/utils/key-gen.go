package utils

import (
	"log"

	"github.com/m-d-nabeel/exploding-kittens/models"
)

func GenerateKeyForRDB(payload interface{}) string {
	var key string
	switch payload := payload.(type) {
	case models.Card:
		key = "card:" + payload.ID.String()
	case models.User:
		key = "user:" + payload.Username
	case models.Game:
		key = "game:" + payload.ID.String()
	default:
		key = ""
	}
	log.Println("Generated key: ", key)
	return key
}
