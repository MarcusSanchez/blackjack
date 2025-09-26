package game

import "slices"

type Player struct {
	money int
	hands []*PlayerHand
}

func NewPlayer(money int) *Player {
	return &Player{
		money: money,
		hands: make([]*PlayerHand, 0, 8), // more than likely won't need more than 8 hands
	}
}

type PlayerHand struct {
	player *Player
	idx    int

	cards    []*Card
	total    int
	bet      int
	standing bool
}

func NewPlayerHand(player *Player, bet int, card1, card2 *Card, splitIdx ...int) *PlayerHand {
	assert(player.money >= bet, "not enough money to place bet")
	player.money -= bet

	hand := &PlayerHand{
		player: player,
		cards:  []*Card{card1, card2},
		bet:    bet,
	}
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
	return ph.total > 21
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
	hand := NewPlayerHand(ph.player, ph.bet, ph.cards[1], card2, ph.idx)
	hand.CalculateTotal()

	ph.cards = []*Card{ph.cards[0], card1}
	ph.CalculateTotal()
}
