// https://www.hackerrank.com/challenges/no-prefix-set/problem

package noprefixset

import (
	"errors"
	"fmt"
	"os"
)

type Trie struct {
	Node   map[rune]*Trie
	IsLeaf bool
}

func NewTrie() *Trie {
	return &Trie{
		Node:   make(map[rune]*Trie),
		IsLeaf: false,
	}
}

func (t *Trie) Add(phrase string) error {
	root := t
	for _, word := range phrase {
		if _, ok := root.Node[word]; !ok {
			root.Node[word] = NewTrie()
		}
		root = root.Node[word]
	}
	root.IsLeaf = true
	return nil
}

func (t *Trie) Search(query string) (*Trie, error) {
	root := t
	for _, word := range query {
		if _, ok := root.Node[word]; !ok {
			return root, errors.New("Not found")
		}
		root = root.Node[word]
	}
	return root, nil
}

func NoPrefixSet(in *os.File) {
	var N int
	trie := NewTrie()
	fmt.Fscanf(in, "%d", &N)
	for n := 0; n < N; n++ {
		var line string
		fmt.Fscanf(in, "%s", &line)
		node, err := trie.Search(line)
		if err != nil { // Not Found!
			if node.IsLeaf {
				fmt.Println("BAD SET")
				fmt.Println(line)
				return
			} else {
				trie.Add(line)
			}
		} else {
			fmt.Println("BAD SET")
			fmt.Println(line)
			return
		}
	}
	fmt.Println("GOOD SET")
}