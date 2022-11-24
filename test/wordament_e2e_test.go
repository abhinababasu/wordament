package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"bonggeek.com/wordament/solver"
	"golang.org/x/exp/slices"
)

const dictionaryPath = "../service/english0.dict"

func TestE2E(t *testing.T) {
	// create and solve the wordament
	input := "SPAVURNYGERSMSBE"
	size := 4
	w := solver.NewWordament(size)
	w.LoadDictionary(dictionaryPath)

	solution, err := w.Solve(input)

	if err != nil {
		t.Error("There should be no error")
	}

	// iterate through all the solutions found
	validWords, err := loadDictionary(dictionaryPath)
	if err != nil {
		t.Fatalf("Load local dict failed with error %v", err)
	}

	longestWordLen := 0
	wordsFound := []string{}
	for _, wordCells := range solution.Result {
		if len(wordCells) > longestWordLen {
			longestWordLen = len(wordCells)
		}

		word := w.WordFromCells(wordCells)
		if _, ok := validWords[word]; !ok {
			t.Errorf("Found word %v which is not in the dictionary", word)
		}
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
	w.LoadDictionary(dictionaryPath)

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

func TestNoDuplicate(t *testing.T) {
	input := "SPAVURNYGERSMSBE"
	size := 4

	w := solver.NewWordament(size)
	w.LoadDictionary(dictionaryPath)

	solution, err := w.Solve(input)

	if err != nil {
		t.Error("There should be no error")
	}

	if len(solution.Result) < 100 {
		t.Errorf("Too few (%v) results", len(solution.Result))
	}

	wordsFound := make(map[string]bool)
	for _, wordCells := range solution.Result {
		word := w.WordFromCells(wordCells)
		_, found := wordsFound[word]
		if found {
			t.Errorf("Duplicate word %v found", word)
		} else {
			wordsFound[word] = true
		}
	}
}

func loadDictionary(path string) (map[string]bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	words := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// remove '
		var word string
		word = strings.ReplaceAll(scanner.Text(), "'", "")
		word = strings.ToUpper(word)
		words[word] = true
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
