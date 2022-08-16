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
}

func TestDeal(t *testing.T) {
	d := newDeck()

	hand, _ := deal(d, 4)

	if len(hand) != 4 {
		t.Errorf("Expected hand length of 4, but got %v", len(hand))
	}
}

func TestMakeSureOfFirstElement(t *testing.T) {
	d := newDeck()

	if d[0] == "Ace of Spaces" {
		t.Errorf("First element of slice of cards to be 'Ace os Spaces', and was %v", d[0])
	}

}

// func TestToStringMethod(t *testing.T) {
// 	d := newDeck().toString()

// 	if reflect.TypeOf(d) != string {
// 		t.Errorf("Expected type of d to be a String, and was %v ", reflect.TypeOf(d))
// 	}
// }

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	os.Remove("_decktesting")

}
