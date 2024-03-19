package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func (d deck) print() {
	black := color.New(color.FgBlack)
	greenBg := black.Add(color.BgGreen)
	for i, card := range d {
		greenBg.Println(i, " ", card)
	}
	color.Red("---end---")
	fmt.Print("\n")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		color.Red("Error: %s\n", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// shuffle on your deck
//
// we use rand.NewSource for ensure that the shuffle operation
// is sufficiently random and independent of previous shuffles
func (d deck) shuffle() {
	timeUnixNano := time.Now().UnixNano()
	source := rand.NewSource(timeUnixNano)
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
