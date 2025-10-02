package game

import (
	"slices"
)

type Player struct {
	money int
	hands []*PlayerHand
}

func (p *Player) Money() int {
	return p.money
}

func NewPlayer(money int) *Player {
	return &Player{
		money: money,
		hands: make([]*PlayerHand, 0, 8), // more than likely won't need more than 8 hands
	}
}

func (p *Player) PrepareHandsForNewRound(bet int, bets ...int) []*PlayerHand {
	assert(max(bet, bets...) <= p.money, "not enough money to place bets")
	p.hands = make([]*PlayerHand, 0, 8)

	bets = slices.Insert(bets, 0, bet)
	for _, b := range bets {
		assert(b > 0, "bet must be greater than 0")
		NewPlayerHand(p, b)
	}

	p.SynchronizeHandIndices()
	return p.hands
}

type PlayerHand struct {
	player *Player
	idx    int

	cards    []*Card
	total    int
	bet      int
	standing bool
}

func NewPlayerHand(player *Player, bet int) *PlayerHand {
	assert(player.money >= bet, "not enough money to place bet")
	player.money -= bet

	hand := &PlayerHand{
		player: player,
		bet:    bet,
	}
	player.hands = append(player.hands, hand)
	player.SynchronizeHandIndices()

	return hand
}

func NewPlayerHandFromCards(player *Player, bet int, card1, card2 *Card, splitIdx ...int) *PlayerHand {
	assert(player.money >= bet, "not enough money to place bet")
	player.money -= bet

	hand := &PlayerHand{
		player: player,
		cards:  []*Card{card1, card2},
		bet:    bet,
	}
	hand.CalculateTotal()

	if splitIdx != nil {
		idx := splitIdx[0] + 1
		assert(idx >= 0 && idx <= len(player.hands), "invalid split index")
		player.hands = slices.Insert(player.hands, idx, hand)
	} else {
		player.hands = append(player.hands, hand)
	}
	player.SynchronizeHandIndices()

	return hand
}

func (p *Player) SynchronizeHandIndices() {
	for i, hand := range p.hands {
		hand.idx = i
	}
}

func (ph *PlayerHand) Busted() bool {
	return ph.total > TotalUpperLimit
}

func (ph *PlayerHand) CalculateTotal() {
	ph.total = calculateHandTotal(ph.cards)
}

func (ph *PlayerHand) Stand() {
	assert(!ph.standing, "hand is already standing")
	assert(len(ph.cards) >= 2, "cannot stand on a hand with less than 2 cards")
	assert(!ph.Busted(), "cannot stand on a busted hand")

	ph.standing = true
}

func (ph *PlayerHand) Hit(card *Card) {
	assert(!ph.standing, "cannot hit on a standing hand")
	assert(!ph.Busted(), "cannot hit on a busted hand")

	ph.cards = append(ph.cards, card)
	ph.CalculateTotal()
}

func (ph *PlayerHand) DoubleDown(card *Card) {
	assert(!ph.standing, "cannot double down on a standing hand")
	assert(len(ph.cards) == 2, "can only double down on a hand with exactly 2 cards")
	assert(ph.player.money >= ph.bet, "not enough money to double down")

	ph.Hit(card)
	ph.bet *= 2
	ph.player.money -= ph.bet
	ph.standing = true
}

func (ph *PlayerHand) Split(card1, card2 *Card) {
	assert(len(ph.cards) == 2, "can only split a hand with exactly 2 cards")
	assert(!ph.Busted(), "cannot split a busted hand")
	assert(ph.cards[0].rank == ph.cards[1].rank, "can only split cards of the same rank")
	assert(ph.player.money >= ph.bet, "not enough money to split")

	ph.player.money -= ph.bet
	hand := NewPlayerHandFromCards(ph.player, ph.bet, ph.cards[1], card2, ph.idx)
	hand.CalculateTotal()

	ph.cards = []*Card{ph.cards[0], card1}
	ph.CalculateTotal()
}
