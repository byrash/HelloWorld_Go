package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_descktesting"
	os.Remove(filename)
	d := newDeck()
	d.saveToFile(filename)
	loadeddeck := newDeckFromFile(filename)
	if len(d) != len(loadeddeck) {
		t.Errorf("Expected %v cards in deck, got %v", len(d), len(loadeddeck))
	}
	os.Remove(filename)
}
