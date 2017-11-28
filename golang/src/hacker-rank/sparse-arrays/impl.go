// https://www.hackerrank.com/challenges/sparse-arrays/problem

package sparsearrays

import (
	"errors"
	"fmt"
	"os"
)

type Trie struct {
	Node     map[rune]*Trie
	Sequence []int
	IsLeaf   bool
}

func NewTrie() *Trie {
	return &Trie{
		Node:     make(map[rune]*Trie),
		Sequence: make([]int, 0),
		IsLeaf:   false,
	}
}

func (t *Trie) Add(phrase string, index int) error {
	root := t
	for _, word := range phrase {
		if _, ok := root.Node[word]; !ok {
			root.Node[word] = NewTrie()
		}
		root = root.Node[word]
	}
	root.Sequence = append(root.Sequence, index)
	root.IsLeaf = true
	return nil
}

func (t *Trie) Search(query string) ([]int, error) {
	root := t
	for _, word := range query {
		if _, ok := root.Node[word]; !ok {
			return nil, errors.New("Not found")
		}
		root = root.Node[word]
	}
	if !root.IsLeaf {
		return nil, errors.New("Does not reach leaf node")
	}
	return root.Sequence, nil
}

func SparseArrays(in *os.File) {
	trie := NewTrie()

	var N, Q int
	fmt.Fscanf(in, "%d", &N)
	for n := 0; n < N; n++ {
		var value string
		fmt.Fscanf(in, "%s", &value)
		trie.Add(value, n)
	}
	fmt.Fscanf(in, "%d", &Q)
	for q := 0; q < Q; q++ {
		var query string
		fmt.Fscanf(in, "%s", &query)
		seq, err := trie.Search(query)
		if err != nil {
			fmt.Println(0)
		} else {
			fmt.Println(len(seq))
		}
	}
}
