package game

import (
	assertgo "github.com/nikoksr/assert-go"
)

var assert = assertgo.Assert

const (
	MaxCardsInHand  = 11 // max 11 cards without busting (generally)
	TotalUpperLimit = 21
	TableSlotsCount = 5 // max 5 starting hands at a table
)

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

func (bj *Blackjack) StartNewRound(bet int, bets ...int) {
	assert(1+len(bets) <= TableSlotsCount, "cannot place more than %d bets", TableSlotsCount)

	hands := bj.player.PrepareHandsForNewRound(bet, bets...)
	bj.dealer.DealRoundOfCards(hands)
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
		if (total + 11) <= TotalUpperLimit {
			total += 11
			continue
		}
		total += 1
	}

	return total
}
