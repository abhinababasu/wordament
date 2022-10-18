package main

import (
	"fmt"
	"strings"
)

type node struct {
	s     string // TODO: Likely this will not be needed
	child map[string]*node
}
type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) AddWord(word string) {
	fmt.Println(word)

	for _, c := range word {
		s := strings.ToUpper(string(c))
		fmt.Print(s, " ")
		// TODO: Now add this stuff to trie
	}
	fmt.Println()
}
