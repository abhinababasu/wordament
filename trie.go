package main

import "fmt"

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
}
