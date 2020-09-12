package main

import (
	"math/rand"
	"sort"
	"time"
)

const (
	Heart   = "Heart"
	Diamond = "Diamond"
	Club    = "Club"
	Spade   = "Spade"
)

var Types = [4]string{
	Heart,
	Diamond,
	Club,
	Spade,
}

const (
	Card1  = "1"
	Card2  = "2"
	Card3  = "3"
	Card4  = "4"
	Card5  = "5"
	Card6  = "6"
	Card7  = "7"
	Card8  = "8"
	Card9  = "9"
	Card10 = "10"
	CardJ  = "J"
	CardQ  = "Q"
	CardK  = "K"
)

var ValueMap = map[string]int{
	Card1:  1,
	Card2:  2,
	Card3:  3,
	Card4:  4,
	Card5:  5,
	Card6:  6,
	Card7:  7,
	Card8:  8,
	Card9:  9,
	Card10: 10,
	CardJ:  11,
	CardQ:  12,
	CardK:  13,
}

var Cards = [13]string{
	Card1,
	Card2,
	Card3,
	Card4,
	Card5,
	Card6,
	Card7,
	Card8,
	Card9,
	Card10,
	CardJ,
	CardQ,
	CardK,
}

type Card struct {
	Type  string
	Value string
}

type Deck struct {
	Cards    []Card
	ValueMap map[string]int
	Draw
	Shuffle
	Comparison
}

type Draw func(deck *Deck, number int) []Card
type Shuffle func(deck *Deck)
type Comparison func(deck *Deck, i, j int) bool

func (d *Deck) WithShuffle(shuffle Shuffle) *Deck {
	d.Shuffle = shuffle
	return d
}

func (d *Deck) WithDraw(draw Draw) *Deck {
	d.Draw = draw
	return d
}

func (d *Deck) WithComparison(comparison Comparison) *Deck {
	d.Comparison = comparison
	return d
}

func (d *Deck) Sort() {
	sort.Sort(d)
}
func RandomShuffle() Shuffle {
	return func(deck *Deck) {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(deck.Cards), func(i, j int) {
			deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
		})
	}
}

func TopDraw() Draw {
	return func(deck *Deck, number int) []Card {
		drawn := deck.Cards[:number]
		deck.Cards = deck.Cards[number:]
		return drawn
	}
}

func StandardComparison() Comparison {
	return func(d *Deck, i, j int) bool {
		return d.ValueMap[d.Cards[i].Value] < d.ValueMap[d.Cards[j].Value]
	}
}

func (d *Deck) Filter(val string) {
	for i, card := range d.Cards {
		if card.Value == val {
			d.Cards = append(d.Cards[:i], d.Cards[i:]...)
		}
	}
}

func NewStandardDeck(count int) *Deck {
	deck := &Deck{
		ValueMap: ValueMap,
	}

	for i := count; i > 0; i-- {
		for _, cardType := range Types {
			for _, name := range Cards {
				deck.Cards = append(deck.Cards, Card{
					Type:  cardType,
					Value: name,
				})
			}
		}
	}

	return deck
}

func (d *Deck) Len() int {
	return len(d.Cards)
}

func (d *Deck) Less(i, j int) bool {
	return d.Comparison(d, i, j)
}

func (d *Deck) Swap(i, j int) {
	d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
}
