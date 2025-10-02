package game

type Dealer struct {
	shoe *Shoe
	hand *DealerHand
}

func NewDealer(decks int) *Dealer {
	dealer := &Dealer{shoe: NewShoe(decks)}
	dealer.hand = NewDealerHand(dealer)
	return dealer
}

func (d *Dealer) DealRoundOfCards(hands []*PlayerHand) {
	assert(len(hands) > 0, "no player hands to deal to")
	assert(len(hands) <= TableSlotsCount, "cannot deal to more than %d hands", TableSlotsCount)

	d.hand = NewDealerHand(d)

	if d.shoe.lastHandPlayed {
		d.shoe = NewShoe(len(d.shoe.cards) / DeckSize)
	}

	for range 2 {
		for _, hand := range hands {
			card := d.shoe.DrawCard()
			hand.Hit(card)
		}

		d.hand.Hit()
	}
}

const (
	UpCardIndex   = 0
	HoleCardIndex = 1
)

type DealerHand struct {
	dealer *Dealer
	cards  []*Card
	total  int
}

func NewDealerHand(dealer *Dealer) *DealerHand {
	return &DealerHand{
		cards:  make([]*Card, 0, MaxCardsInHand),
		total:  0,
		dealer: dealer,
	}
}

func (dh *DealerHand) Busted() bool {
	return dh.total > TotalUpperLimit
}

func (dh *DealerHand) CalculateTotal() {
	dh.total = calculateHandTotal(dh.cards)
}

func (dh *DealerHand) UpCard() *Card {
	assert(len(dh.cards) < 1, "dealer has no cards")
	return dh.cards[UpCardIndex]
}

func (dh *DealerHand) HoleCard() *Card {
	assert(len(dh.cards) < 1, "dealer has no cards")
	assert(len(dh.cards) < HoleCardIndex+1, "dealer has no hole card")
	return dh.cards[HoleCardIndex]
}

func (dh *DealerHand) CheckForBlackjack() bool {
	assert(len(dh.cards) == 2, "dealer must have exactly 2 cards to check for blackjack")
	assert(dh.UpCard().rank >= Ten, "dealer's up card must be 10, J, Q, K, or A to check for blackjack")
	return calculateHandTotal(dh.cards) == TotalUpperLimit
}

func (dh *DealerHand) PlayOutHand() {
	assert(len(dh.cards) == 2, "dealer must have exactly 2 cards to play out hand")
	assert(!dh.Busted(), "dealer cannot play out a busted hand")

	for dh.total < 17 {
		dh.Hit()
	}
}

func (dh *DealerHand) Hit() {
	card := dh.dealer.shoe.DrawCard()
	dh.cards = append(dh.cards, card)
	dh.CalculateTotal()
}
