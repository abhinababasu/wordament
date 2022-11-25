package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"bonggeek.com/wordament/solver"
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

	expectedWords := []string{
		"SURGERY",
		"SPARES",
	}
	longestWordLen, err := validateCorrectWordsFound(*w, solution, expectedWords)
	if err != nil {
		t.Fatalf("Tests failed '%v'", err)
	}
	expectedLongestLen := 7
	if longestWordLen != expectedLongestLen {
		t.Errorf("The longest word found should be %v character long", longestWordLen)
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

var wg sync.WaitGroup
var result1 []string
var result2 []string

func TestParallelSolve(t *testing.T) {
	// create and solve the wordament
	input1 := "SPAVURNYGERSMSBE"
	input2 := "ZRFLPFUALINXAYEM"
	size := 4
	w := solver.NewWordament(size)
	w.LoadDictionary(dictionaryPath)

	results := make(chan solver.WordamentResult)
	solveFunc := func(input string) {
		solution, _ := w.Solve(input)
		results <- solution
	}

	go solveFunc(input1)
	go solveFunc(input2)

	fmt.Println("Waiting for results")
	results1 := <-results
	results2 := <-results

	// the results
	fmt.Println("Result 1: !!!", getInputFromMatrix(results1.Input), results1)
	fmt.Println("Result 2: !!!", getInputFromMatrix(results2.Input), results2)

	// TODO: actually validate the results are not being mixed up (there is a bug and it is currently)
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

func getInputFromMatrix(m [][]rune) string {
	s := ""
	for i, _ := range m {
		for j, _ := range m[i] {
			s = s + string(m[i][j])
		}
	}

	return s
}

func validateCorrectWordsFound(w solver.Wordament, solution solver.WordamentResult, expectedWords []string) (int, error) {
	// iterate through all the solutions found
	validWords, err := loadDictionary(dictionaryPath)
	if err != nil {
		return 0, fmt.Errorf("Load local dict failed with error %v", err)
	}

	if len(solution.Result) < 50 {
		return 0, fmt.Errorf("Too few words found")
	}
	longestWordLen := 0
	wordsFound := map[string]bool{}
	for _, wordCells := range solution.Result {
		if len(wordCells) > longestWordLen {
			longestWordLen = len(wordCells)
		}

		word := w.WordFromCells(wordCells)
		if _, ok := validWords[word]; !ok {
			return 0, fmt.Errorf("Found word %v which is not in the dictionary", word)
		}
		wordsFound[word] = true
	}

	for _, expectedWord := range expectedWords {
		if _, ok := wordsFound[expectedWord]; !ok {
			return 0, fmt.Errorf("Expected word %v not found", expectedWord)
		}
	}

	return longestWordLen, nil
}
