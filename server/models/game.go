package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type GameStatus string

const (
	ActiveGame   GameStatus = "Active"
	FinishedGame GameStatus = "Finished"
)

type Game struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	UserId    uuid.UUID  `json:"user_id"`
	Deck      [5]Card    `json:"deck"`
	Status    GameStatus `json:"status"`
	Score     int        `json:"score"`
}

func getRandomCards() [5]Card {
	deck := [5]Card{}
	size := 5
	for i := 0; i < size; i++ {
		rndIdx := rand.Intn(len(CardTypes))
		deck[i] = NewCard(CardTypes[rndIdx])
	}
	return deck
}

func NewGame(userId uuid.UUID) Game {
	deck := getRandomCards()
	return Game{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserId:    userId,
		Status:    ActiveGame,
		Score:     0,
		Deck:      deck,
	}
}
