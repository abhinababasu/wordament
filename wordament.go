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
	trie *Trie
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

	w.trie = NewTrie()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// remove '
		var word string
		word = strings.ReplaceAll(scanner.Text(), "'", "")
		w.trie.AddWord(word)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	word := "inscribe"
	fmt.Println(word, w.trie.WordExists(word))

	word = "inscribed"
	fmt.Println(word, w.trie.WordExists(word))

	word = "inscr"
	fmt.Println(word, w.trie.WordExists(word))

	word = "abhinaba"
	fmt.Println(word, w.trie.WordExists(word))

}

func (w *Wordament) Solve(matrix [][]string) [][]int {
	return nil
}
