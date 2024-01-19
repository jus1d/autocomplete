package trie

import (
	"strings"
)

type Trie struct {
	root *node
}

type node struct {
	children     map[rune]*node
	isTerminated bool
}

func New() *Trie {
	node := node{
		children:     make(map[rune]*node),
		isTerminated: false,
	}

	return &Trie{
		root: &node,
	}
}

func (t *Trie) Autocomplete(substring string) []string {
	substring = strings.ToLower(substring)

	cur := t.root
	for _, ch := range substring {
		if cur.children[ch] == nil {
			return []string{}
		}
		cur = cur.children[ch]
	}

	return cur.getWords(substring)
}

func (n *node) getWords(sub string) []string {
	cur := n
	words := make([]string, 0)

	for k, v := range cur.children {
		w := v.getWords(sub + string(k))
		for i := 0; i < len(w); i++ {
			words = append(words, w[i])
		}
	}

	if cur.isTerminated {
		words = append(words, sub)
	}

	return words
}

func (n *node) countKeys() int {
	count := 0
	for range n.children {
		count++
	}
	return count
}

func (t *Trie) Insert(word string) {
	word = strings.ToLower(word)

	cur := t.root
	for _, ch := range word {
		if cur.children[ch] == nil {
			cur.children[ch] = &node{
				children:     make(map[rune]*node),
				isTerminated: false,
			}
		}
		cur = cur.children[ch]
	}
	cur.isTerminated = true
}

func (t *Trie) InsertMany(words []string) {
	for _, word := range words {
		t.Insert(word)
	}
}

func (t *Trie) ContainsWord(word string) bool {
	word = strings.ToLower(word)

	cur := t.root
	for _, ch := range word {
		if cur.children[ch] == nil {
			return false
		}
		cur = cur.children[ch]
	}
	return cur.isTerminated
}

func (t *Trie) ContainsSubstring(substring string) bool {
	substring = strings.ToLower(substring)

	cur := t.root
	for _, ch := range substring {
		if cur.children[ch] == nil {
			return false
		}
		cur = cur.children[ch]
	}
	return true
}
