package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const limit = 10

func main() {
	var n int
	var arr []string
	fmt.Scanf("%d", &n)
	arr = make([]string, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		if len(s) > limit {
			var buf bytes.Buffer
			buf.WriteByte(s[0])
			ins := len(s[1 : len(s)-1])
			buf.WriteString(strconv.Itoa(ins))
			buf.WriteByte(s[len(s)-1])
			arr = append(arr, buf.String())
		} else {
			arr = append(arr, s)
		}
	}
	fmt.Println(strings.Join(arr, "\n"))
}
