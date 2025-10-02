package game

type Suit int

const (
	Hearts Suit = iota + 1
	Diamonds
	Clubs
	Spades
)

func (s Suit) String() string {
	switch s {
	case Hearts:
		return "Hearts"
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	case Spades:
		return "Spades"
	}
	return "Unknown"
}

type Rank int

const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (v Rank) String() string {
	switch v {
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	}
	return "Unknown"
}

func (v Rank) Value() int {
	if v >= Two && v <= Ten {
		return int(v)
	}
	if v >= Jack && v <= King {
		return 10
	}
	if v == Ace {
		assert(false, "ace value is either 1 or 11, cannot determine single value") // dump stack for debugging
	}
	return 0
}
