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

type Shoe []Card

func NewShoe(decks int) Shoe {
	shoe := make(Shoe, 0, decks*DeckSize)
	for range decks {
		cards := NewDeck()
		shoe = append(shoe, cards...)
	}

	low, high := CutCardLowerBound(decks), CutCardUpperBound(decks)
	cut := low + rand.Intn(high-low+1)
	shoe[cut].cut = true

	shoe.Shuffle()
	return shoe
}

func NewDeck() []Card {
	cards := make([]Card, 0, 52)
	for suit := Hearts; suit <= Spades; suit++ {
		for rank := Two; rank <= Ace; rank++ {
			card := Card{suit: suit, rank: rank}
			cards = append(cards, card)
		}
	}
	return cards
}

func (s Shoe) Shuffle() {
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
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
