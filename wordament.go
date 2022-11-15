package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Wordament struct {
	size   int
	trie   *Trie
	matrix [][]rune
	result []string
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

func (w *Wordament) Solve(matrix [][]rune) []string {
	cells := []Cell{}
	w.matrix = matrix
	w.result = []string{}
	for r := 0; r < w.size; r++ {
		for c := 0; c < w.size; c++ {

			ch := w.matrix[r][c]
			node := w.trie.root.GetChild(ch)

			if node != nil {
				w.solvePos(GetCell(r, c), "", node, cells)
			}
		}
	}
	return w.result
}

func (w *Wordament) solvePos(cell Cell, s string, trn *node, cells []Cell) {
	newWord := s + trn.s
	if trn.IsWordEnd() {
		w.result = append(w.result, newWord)
	}

	ncells := cell.GetNeighbors(w.size-1, w.size-1)
	newList := append(cells, cell)
	for _, ncell := range ncells {
		if ncell.CellInList(cells) {
			continue
		}
		r, c := ncell.row, ncell.col
		ch := w.matrix[r][c]
		child := trn.GetChild(ch)
		if child == nil {
			continue
		}

		w.solvePos(ncell, newWord, child, newList)
	}
}
