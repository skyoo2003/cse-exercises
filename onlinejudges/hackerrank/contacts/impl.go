package contacts

import (
	"errors"
	"fmt"
	"os"
)

type Trie struct {
	Node     map[rune]*Trie
	Contacts []int
	IsLeaf   bool
}

func NewTrie() *Trie {
	return &Trie{
		Node:     make(map[rune]*Trie),
		Contacts: make([]int, 0),
		IsLeaf:   false,
	}
}

func (t *Trie) Add(phrase string, no int) error {
	root := t
	for _, word := range phrase {
		if _, ok := root.Node[word]; !ok {
			root.Node[word] = NewTrie()
		}
		root.Contacts = append(root.Contacts, no)
		root = root.Node[word]
	}
	root.Contacts = append(root.Contacts, no)
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
	return root.Contacts, nil
}

func Contact(in *os.File) {
	var N int
	fmt.Fscanf(in, "%d", &N)

	trie := NewTrie()

	for n := 0; n < N; n++ {
		var cmd, arg string
		fmt.Fscanf(in, "%s %s", &cmd, &arg)
		switch cmd {
		case "add":
			trie.Add(arg, n)
		case "find":
			contacts, err := trie.Search(arg)
			if err != nil {
				fmt.Println(0)
			} else {
				fmt.Println(len(contacts))
			}
		}
	}
}
