package game

import "math/rand"

const (
	DeckSize = 52

	singleDeckCutCardLowerBound = DeckSize / 2
	multiDeckCutCardLowerBound  = DeckSize * 1.5
)

type Card struct {
	suit Suit
	rank Rank

	// cut is a special marker used to indicate when the last hand is being played
	cut bool
}

type Shoe struct {
	cards          []*Card
	lastHandPlayed bool
}

func NewShoe(decks int) *Shoe {
	shoe := &Shoe{cards: make([]*Card, 0, decks*DeckSize)}
	for range decks {
		cards := NewDeck()
		shoe.cards = append(shoe.cards, cards...)
	}
	shoe.Shuffle()

	low, high := CutCardLowerBound(decks), CutCardUpperBound(decks)
	cut := low + rand.Intn(high-low+1)
	shoe.cards[cut].cut = true

	return shoe
}

func (s *Shoe) Shuffle() {
	for i := len(s.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s.cards[i], s.cards[j] = s.cards[j], s.cards[i]
	}
}

func (s *Shoe) DrawCard() *Card {
	assert(len(s.cards) > 0, "shoe is empty, cannot draw card")

	card := s.cards[0]
	s.cards = s.cards[1:]

	if card.cut {
		s.lastHandPlayed = true
	}

	return card
}

func NewDeck() []*Card {
	cards := make([]*Card, 0, 52)
	for suit := Hearts; suit <= Spades; suit++ {
		for rank := Two; rank <= Ace; rank++ {
			card := &Card{suit: suit, rank: rank}
			cards = append(cards, card)
		}
	}
	return cards
}

func CutCardLowerBound(decks int) int {
	if decks <= 1 {
		return singleDeckCutCardLowerBound
	}
	return multiDeckCutCardLowerBound
}

func CutCardUpperBound(decks int) int {
	return (decks * DeckSize) - (DeckSize * 1.5)
}
