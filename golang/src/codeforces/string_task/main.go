package main

import (
	"bytes"
	"fmt"
	"strings"
)

var Vowels = map[byte]bool{
	'a': true,
	'o': true,
	'y': true,
	'e': true,
	'u': true,
	'i': true,
}

func main() {
	var s string
	var buf bytes.Buffer
	fmt.Scan(&s)
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if _, ok := Vowels[ch]; ok {
			continue
		} else {
			buf.WriteByte('.')
			buf.WriteByte(ch)
		}
	}
	fmt.Println(buf.String())
}
