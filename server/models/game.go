package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type GameStatus string

type GameMoveResult string

const (
	SuccessMove GameMoveResult = "Success"
	DiffuseMove GameMoveResult = "Diffuse"
	ExplodeMove GameMoveResult = "Explode"
	ShuffleMove GameMoveResult = "Shuffle"
	NoMove      GameMoveResult = "NoMove"
)

const (
	ActiveGame   GameStatus = "Active"
	FinishedGame GameStatus = "Finished"
)

type Game struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UserId      uuid.UUID  `json:"user_id"`
	Deck        []Card     `json:"deck"`
	DiffuseCard int        `json:"diffuse_card"`
	Status      GameStatus `json:"status"`
	Score       int        `json:"score"`
}

func GetRandomCards() []Card {
	deck := []Card{}
	size := 5
	for i := 0; i < size; i++ {
		rndIdx := rand.Intn(len(CardTypes))
		deck = append(deck, NewCard(CardTypes[rndIdx]))
	}
	return deck
}

func NewGame(userId uuid.UUID, status GameStatus) Game {
	deck := GetRandomCards()
	return Game{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		UserId:      userId,
		DiffuseCard: 0,
		Status:      status,
		Score:       0,
		Deck:        deck,
	}
}
