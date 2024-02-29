package utils

import "github.com/m-d-nabeel/exploding-kittens/models"

func GetGameResultForMove(game *models.Game, cardType *models.CardType) (*models.Game, error) {
	switch *cardType {
	case models.DiffuseCard:
		game.DiffuseCard++
		game.Score += 2
		game.Deck = [5]models.Card(removeCardFromDeck(game.Deck, models.Card{Type: models.DiffuseCard}))
	case models.ExplodingCard:
		if game.DiffuseCard > 0 {
			game.DiffuseCard--
			game.Score += 5
			game.Deck = [5]models.Card(removeCardFromDeck(game.Deck, models.Card{Type: models.ExplodingCard}))
		} else {
			game.Status = models.FinishedGame
		}
	case models.ShuffleCard:
		game.Deck = models.GetRandomCards()
	case models.CatCard:
		game.Score += 1
		game.Deck = [5]models.Card(removeCardFromDeck(game.Deck, models.Card{Type: models.CatCard}))
	}
	return nil, nil
}

func removeCardFromDeck(deck [5]models.Card, card models.Card) [5]models.Card {
	for _, c := range deck {
		if c.ID == card.ID {
			c.Type = models.EmptyCard
			break
		}
	}
	return deck
}
