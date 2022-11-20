package main

import (
	"testing"

	"bonggeek.com/wordament/solver"
	"golang.org/x/exp/slices"
)

func TestE2E(t *testing.T) {

	input := "SPAVURNYGERSMSBE"
	size := 4

	w := solver.NewWordament(size)
	w.LoadDictionary("../service/english0.dict")

	solution, err := w.Solve(input)

	if err != nil {
		t.Error("There should be no error")
	}

	longestWordLen := 0
	wordsFound := []string{}
	for _, wordCells := range solution.Result {
		if len(wordCells) > longestWordLen {
			longestWordLen = len(wordCells)
		}

		word := w.WordFromCells(wordCells)
		wordsFound = append(wordsFound, word)
	}

	expectedLongestLen := 7
	if longestWordLen != expectedLongestLen {
		t.Errorf("The longest word found should be %v character long", longestWordLen)
	}

	idx := slices.IndexFunc(wordsFound, func(s string) bool { return s == "SURGERY" })
	if idx == -1 {
		t.Error("Word not found")
	}
}

func TestBadInput(t *testing.T) {

	size := 4

	w := solver.NewWordament(size)
	w.LoadDictionary("../solver/english.0")
	w.LoadDictionary("../solver/english.2")

	input := "SPAVURN"
	_, err := w.Solve(input)
	if err == nil {
		t.Error("Should fail as input too short")
	}

	input = "SPAVURNWERWERWEREWR"
	_, err = w.Solve(input)
	if err == nil {
		t.Error("Should fail as input too long")
	}
}
