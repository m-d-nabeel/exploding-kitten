package models

import (
	"time"

	"github.com/google/uuid"
)

type CardType string

const (
	CatCard       CardType = "Cat"
	DiffuseCard   CardType = "Diffuse"
	ShuffleCard   CardType = "Shuffle"
	ExplodingCard CardType = "Exploding"
	EmptyCard     CardType = ""
)

var CardTypes = []CardType{CatCard, DiffuseCard, ShuffleCard, ExplodingCard}

type Card struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Type      CardType  `json:"type"`
}

func NewCard(cardType CardType) Card {
	return Card{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Type:      cardType,
	}
}
