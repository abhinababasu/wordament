package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Wordament struct {
	size int
}

func NewWordament(sz int) *Wordament {
	return &Wordament{size: sz}
}

func (w *Wordament) LoadDictionary(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// remove '
		var word string
		word = strings.ReplaceAll(scanner.Text(), "'", "")

		// TODO: Replace following by a trie add word
		fmt.Println(word)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (w *Wordament) Solve(matrix [][]string) [][]int {
	return nil
}
