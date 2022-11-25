package solver

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var trie *Trie // Single instance of trie

// All dictionaries needs to be loaded at the start
// Note that this is NOT thread-safe
func LoadDictionary(path string) error {
	if trie == nil {
		trie = NewTrie()
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// remove '
		var word string
		word = strings.ReplaceAll(scanner.Text(), "'", "")
		trie.AddWord(word)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

type Wordament struct {
	size int

	matrix [][]rune
	result [][]Cell
}

type WordamentResult struct {
	Input  [][]rune
	Result [][]Cell
}

func NewWordament(sz int) *Wordament {
	// trie needs to be initialized before creating Wordament
	if trie == nil {
		return nil
	}

	w := Wordament{size: sz}
	return &w
}

func (w *Wordament) Solve(s string) (WordamentResult, error) {
	if len(s) != w.size*w.size {
		return WordamentResult{}, fmt.Errorf("Invalid input string. It should be of length %v", w.size*w.size)
	}

	m := w.parseMatrix(s)

	solution := w.SolveMatrix(m)

	result := WordamentResult{
		Input:  m,
		Result: solution,
	}

	return result, nil

}

func (w *Wordament) SolveMatrix(matrix [][]rune) [][]Cell {
	cells := []Cell{}
	w.matrix = matrix
	w.result = [][]Cell{}
	wordsFound := make(map[string]bool)
	for r := 0; r < w.size; r++ {
		for c := 0; c < w.size; c++ {
			ch := w.matrix[r][c]
			node := trie.root.GetChild(ch)

			if node != nil {
				w.solvePos(wordsFound, GetCell(r, c), node, cells)
			}
		}
	}

	// Sort the results with the descending sizes because longer words are more points
	sort.Slice(w.result, func(i, j int) bool {
		return len(w.result[i]) > len(w.result[j])
	})

	return w.result
}

func (w *Wordament) solvePos(wordsFound map[string]bool, cell Cell, trn *node, cells []Cell) {
	// recursively find the words starting at any position in the matrix
	// cell is the current cell
	// trn is the trie node being pointed to currently
	// cells is all the cells in the current discovery path

	// if  we are already in a end of word, add it to solution
	if trn.IsWordEnd() {
		currCells := make([]Cell, len(cells)+1)
		copy(currCells, cells)
		currCells[len(currCells)-1] = cell
		currWord := w.WordFromCells(currCells)
		if _, found := wordsFound[currWord]; !found {
			wordsFound[currWord] = true
			w.result = append(w.result, currCells)
		}
	}

	ncells := cell.GetNeighbors(w.size-1, w.size-1)
	newList := append(cells, cell)
	for _, ncell := range ncells {
		if ncell.CellInList(cells) {
			continue
		}
		r, c := ncell.Row, ncell.Col
		ch := w.matrix[r][c]
		child := trn.GetChild(ch)
		if child == nil {
			continue
		}

		w.solvePos(wordsFound, ncell, child, newList)
	}
}

func (w *Wordament) WordFromCells(cells []Cell) string {
	runes := []rune{}
	for _, wc := range cells {
		runes = append(runes, w.matrix[wc.Row][wc.Col])
	}

	return string(runes)
}

func (w *Wordament) parseMatrix(s string) [][]rune {
	matrix := make([][]rune, w.size)
	for i := range matrix {
		matrix[i] = make([]rune, w.size)
	}

	// we are just assuming these to be english chars
	k := 0
	for i := 0; i < w.size; i++ {
		for j := 0; j < w.size; j++ {
			ch := s[k]
			k++
			matrix[i][j] = rune(ch)
		}
	}

	return matrix
}
