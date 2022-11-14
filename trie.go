package main

import (
	"strings"
)

type node struct {
	s string // TODO: Likely this will not be needed

	child map[rune]*node
}
type Trie struct {
	root *node
}

func NewTrie() *Trie {
	tr := Trie{}
	tr.root = getNode()
	return &tr
}

func (t *Trie) AddWord(word string) {
	p := t.root
	word = strings.ToUpper(word)

	// iterate through the characters in the word
	for _, c := range word {
		// starting at root at every step we see if the current trie node
		// has a child with that same char. If yes we move to that node
		// if not we add a node for that char and do the same
		np, ok := p.child[c]
		if !ok {
			np = getNode()
			p.child[c] = np
		}

		p = np
	}

	// marking end of a word by add a null rune node
	if _, ok := p.child[0]; !ok {
		p.child[0] = getNode()
	}

}

func (t *Trie) WordExists(word string) bool {
	p := t.root
	word = strings.ToUpper(word)

	// iterate through the characters in the word
	for _, c := range word {
		// starting at root at every step we see if the current trie node
		// has a child with the current char. If not the word is not found
		// if found we move to that node
		np, ok := p.child[c]
		if !ok {
			return false
		}

		p = np
	}

	// we have found all the chars of the word in the trie. But the final
	// node we have found also has to be marked as end of an word
	// that is done by a 0/null char, verify that exists
	_, ok := p.child[0]
	return ok

}

func getNode() *node {
	n := node{}
	n.child = make(map[rune]*node)

	return &n
}

func (n *node) GetChild(ch rune) *node {
	if n == nil {
		return nil
	}

	if np, ok := n.child[ch]; ok {
		return np
	} else {
		return nil
	}

}
