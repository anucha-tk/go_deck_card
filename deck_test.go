package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 9 {
		t.Errorf("Expected deck length of 9, but got %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	const name = "_decktesting"
	os.Remove(name)

	d := newDeck()
	d.saveToFile(name)

	loadDeck := newDeckFromFile(name)
	if len(loadDeck) != 9 {
		t.Errorf("Expected 9 cards in deck %v", len(loadDeck))
	}

	os.Remove(name)
}
