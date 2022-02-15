package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 6 {
		t.Errorf("Expected deck length of 6, but got %v", len(d))
	}
}

func TestSaveToFile(t *testing.T) {
	filename := "_decktetsing"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := readFromFile(filename)

	if len(loadedDeck) != 6 {
		t.Errorf("Expected deck length of 6, but got %v", len(d))
	}

	os.Remove(filename)
}
