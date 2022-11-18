package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

type Wordament struct {
	size   int
	trie   *Trie
	matrix [][]rune
	result [][]Cell
}

func NewWordament(sz int) *Wordament {
	w := Wordament{size: sz}
	w.trie = NewTrie()

	return &w
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
		w.trie.AddWord(word)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func (w *Wordament) Solve(matrix [][]rune) [][]Cell {
	cells := []Cell{}
	w.matrix = matrix
	w.result = [][]Cell{}
	for r := 0; r < w.size; r++ {
		for c := 0; c < w.size; c++ {
			ch := w.matrix[r][c]
			node := w.trie.root.GetChild(ch)

			if node != nil {
				w.solvePos(GetCell(r, c), node, cells)
			}
		}
	}

	// Sort the results with the descending sizes because longer words are more points
	sort.Slice(w.result, func(i, j int) bool {
		return len(w.result[i]) > len(w.result[j])
	})

	return w.result
}

func (w *Wordament) solvePos(cell Cell, trn *node, cells []Cell) {
	// recursively find the words starting at any position in the matrix
	// cell is the current cell
	// trn is the trie node being pointed to currently
	// cells is all the cells in the current discovery path

	// if  we are already in a end of word, add it to solution
	if trn.IsWordEnd() {
		currCells := make([]Cell, len(cells)+1)
		copy(currCells, cells)
		currCells[len(currCells)-1] = cell
		w.result = append(w.result, currCells)
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

		w.solvePos(ncell, child, newList)
	}
}

func (w *Wordament) WordFromCells(cells []Cell) string {
	runes := []rune{}
	for _, wc := range cells {
		runes = append(runes, w.matrix[wc.row][wc.col])
	}

	return string(runes)
}
