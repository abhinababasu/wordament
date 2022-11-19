package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestE2E(t *testing.T) {

	input := "SPAVURNYGERSMSBE"

	w := NewWordament(Size)
	w.LoadDictionary("english.0")
	w.LoadDictionary("english.2")

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
