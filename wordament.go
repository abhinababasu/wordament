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
	r, c := 1, 1
	ch := []rune(matrix[r][c])
	node := w.trie.root.GetChild(ch[0])
	if node == nil {
		return nil
	}
	w.solvePos(matrix, r, c, "", node, cells)
	return nil
}

func (w *Wordament) solvePos(matrix [][]string, r, c int, s string, trn *node, cells []Cell) {
	cell := GetCell(r, c)
	newWord := s + trn.s
	if trn.IsWordEnd() {
		fmt.Println(newWord) // this is a word
	}

	ncells := cell.GetNeighbors(w.size-1, w.size-1)
	newList := append(cells, cell)
	for _, ncell := range ncells {
		if ncell.CellInList(cells) {
			continue
		}
		r, c := ncell.row, ncell.col
		ch := []rune(matrix[r][c])
		child := trn.GetChild(ch[0])
		if child == nil {
			continue
		}

		w.solvePos(matrix, ncell.row, ncell.col, newWord, child, newList)
	}
}
