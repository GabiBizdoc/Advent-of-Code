package solution

import (
	"fmt"
	"slices"
)

type cardValue int

var cardsMap map[rune]cardValue

func init() {
	cardsMap = make(map[rune]cardValue)
	for i, card := range "23456789TJQKA" {
		cardsMap[card] = cardValue(i + 1)
	}
}

type HandType int

const (
	Invalid = iota
	FiveOfKind
	FourOfKind
	FullHouse
	ThreeOfKind
	TwoParis
	OnePair
	HighCard
)

type Hand struct {
	o           string
	Cards       [5]cardValue
	orderedHand []cardValue
	kind        HandType
	Bid         int
}

func NewHand(hand string, bid int) *Hand {
	h := &Hand{Bid: bid}
	h.SetHand(hand)
	h.o = hand
	return h
}

func (h *Hand) SetHand(hand string) {
	for i, c := range hand {
		h.Cards[i] = cardsMap[c]
	}
	h.updateType()
}

func (h *Hand) updateType() {
	m := make([]cardValue, 20)
	for _, card := range h.Cards {
		m[card] += 1
	}
	slices.SortFunc(m, func(a, b cardValue) int {
		return int(b - a)
	})
	h.orderedHand = m[:5]

	switch m[0] {
	case 5:
		h.kind = FiveOfKind
	case 4:
		h.kind = FourOfKind
	case 3:
		switch m[1] {
		case 2:
			h.kind = FullHouse
		case 1:
			h.kind = ThreeOfKind
		default:
			h.kind = Invalid
		}
	case 2:
		switch m[1] {
		case 2:
			h.kind = TwoParis
		case 1:
			h.kind = OnePair
		default:
			h.kind = Invalid
		}
	case 1:
		h.kind = HighCard
	default:
		h.kind = Invalid
	}

	if h.kind == Invalid {
		fmt.Printf("%-v\n", h)
		panic("invalid kind")
	}
}

func CmpHands(a *Hand, b *Hand) int {
	if a.kind != b.kind {
		return int(a.kind - b.kind)
	}
	for i, aCard := range a.Cards {
		bCard := b.Cards[i]
		if aCard != bCard {
			return int(bCard - aCard)
		}
	}

	return 0
}
