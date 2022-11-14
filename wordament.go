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

}

func (w *Wordament) Solve(matrix [][]string) [][]int {
	cells := []Cell{}
	w.solvePos(2, 1, "", w.trie.root, cells)
	return nil
}

func (w *Wordament) solvePos(r, c int, s string, trn *node, cells []Cell) {
	cell := GetCell(r, c)
	nc := cell.GetNeighbors(w.size-1, w.size-1)
	fmt.Println("neighbors", nc)
}
