package game

import assertgo "github.com/nikoksr/assert-go"

var assert = assertgo.Assert

type Blackjack struct {
	dealer *Dealer
	player *Player
}

func NewBlackjack(decks int, money int) *Blackjack {
	return &Blackjack{
		dealer: NewDealer(decks),
		player: NewPlayer(money),
	}
}

func calculateHandTotal(cards []*Card) int {
	total := 0
	aces := 0

	for _, card := range cards {
		if card.rank == Ace {
			aces++
			continue
		}
		total += card.rank.Value()
	}

	for range aces {
		if (total + 11) <= 21 {
			total += 11
			continue
		}
		total += 1
	}

	return total
}
