package game

type Dealer struct {
	shoe Shoe
	hand *DealerHand
}

func NewDealer(decks int) *Dealer {
	return &Dealer{
		shoe: NewShoe(decks),
		hand: NewDealerHand(),
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

func NewDealerHand() *DealerHand {
	return &DealerHand{
		cards: make([]*Card, 0, 11), // max 11 cards without busting (A,2,3,4,5,6)
		total: 0,
	}
}

func (dh *DealerHand) Busted() bool {
	return dh.total > 21
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
	return calculateHandTotal(dh.cards) == 21
}

func (dh *DealerHand) PlayOutHand() {
	assert(len(dh.dealer.shoe) > 0, "shoe is empty, cannot play out hand")
	assert(len(dh.cards) == 2, "dealer must have exactly 2 cards to play out hand")
	assert(!dh.Busted(), "dealer cannot play out a busted hand")

	for dh.total < 17 {
		dh.hit()
	}
}

func (dh *DealerHand) hit() {
	assert(len(dh.dealer.shoe) > 0, "shoe is empty, cannot hit")

	card := dh.dealer.shoe[0]
	dh.dealer.shoe = dh.dealer.shoe[1:]
	dh.cards = append(dh.cards, &card)
	dh.CalculateTotal()
}
