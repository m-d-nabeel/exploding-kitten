package utils

import (
	"github.com/m-d-nabeel/exploding-kittens/models"
)

func GetGameResultForMove(game *models.Game, cardType models.CardType) (*models.Game, error) {
	if game.Status == models.FinishedGame {
		return game, nil
	}
	switch cardType {
	case models.DiffuseCard:
		game.DiffuseCard++
		game.Score += 2
		removeCardFromDeck(&game.Deck, models.DiffuseCard)
	case models.ExplodingCard:
		if game.DiffuseCard > 0 {
			game.DiffuseCard--
			game.Score += 5
			removeCardFromDeck(&game.Deck, models.ExplodingCard)
		} else {
			game.Status = models.FinishedGame
		}
	case models.ShuffleCard:
		game.Deck = models.GetRandomCards()
	case models.CatCard:
		game.Score += 1
		removeCardFromDeck(&game.Deck, models.CatCard)
	}
	return game, nil
}

func removeCardFromDeck(deck *[]models.Card, cardToRemoveType models.CardType) {
	for i := range len(*deck) {
		if (*deck)[i].Type == cardToRemoveType {
			(*deck)[i] = models.NewCard(models.EmptyCard)
			break
		}
	}
}
