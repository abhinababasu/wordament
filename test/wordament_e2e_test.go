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
	solver.LoadDictionary(dictionaryPath)
	w := solver.NewWordament(size)

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

	solver.LoadDictionary(dictionaryPath)
	w := solver.NewWordament(size)

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

	solver.LoadDictionary(dictionaryPath)
	w := solver.NewWordament(size)

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
	// create  the wordament
	input1 := "SPAVURNYGERSMSBE"
	input2 := "ZRFLPFUALINXAYEM"
	size := 4
	solver.LoadDictionary(dictionaryPath)
	wordament1 := solver.NewWordament(size)
	wordament2 := solver.NewWordament(size)

	results := make(chan solver.WordamentResult, 2)
	solveFunc := func(w *solver.Wordament, input string) {
		solution, _ := w.Solve(input)
		results <- solution
	}

	// TODO: solve in parallel
	go solveFunc(wordament1, input1)
	go solveFunc(wordament2, input2)

	fmt.Println("Waiting for results")
	results1 := <-results
	results2 := <-results

	resultString1 := getInputFromMatrix(results1.Input)
	resultString2 := getInputFromMatrix(results2.Input)

	if strings.Compare(resultString1, resultString2) == 0 {
		t.Errorf("Both results cannot have the same input %v", resultString1)
		return
	}

	// results coulld've come in any order, so lets swap if necessary
	if resultString1 != input1 {
		resultString1, resultString2 = resultString2, resultString1
		results1, results2 = results2, results1
	}

	expectedWords1 := []string{
		"SURGERY",
		"SPARES",
	}
	_, err1 := validateCorrectWordsFound(*wordament1, results1, expectedWords1)
	if err1 != nil {
		t.Errorf("Validation for %v failed with '%v", resultString1, err1)
	}

	expectedWords2 := []string{
		"ALPINE",
		"MENIAL",
	}
	_, err2 := validateCorrectWordsFound(*wordament2, results2, expectedWords2)
	if err2 != nil {
		t.Errorf("Validation for %v failed with '%v", resultString2, err2)
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

	longestWordLen := 0
	wordsFound := map[string]bool{}
	for _, wordCells := range solution.Result {
		if len(wordCells) > longestWordLen {
			longestWordLen = len(wordCells)
		}

		word := w.WordFromCells(wordCells)

		//fmt.Println("Found ", word)
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

	if len(solution.Result) < 30 {
		return 0, fmt.Errorf("Too few (%v) words found", len(solution.Result))
	}

	return longestWordLen, nil
}
