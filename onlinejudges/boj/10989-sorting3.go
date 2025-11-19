package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	var counter []int

	fmt.Fscan(reader, &N)
	counter = make([]int, 10000)
	for i := 0; i < N; i++ {
		var value int
		fmt.Fscan(reader, &value)
		counter[value-1] += 1
	}

	for i := 0; i < len(counter); i++ {
		count := counter[i]
		if count == 0 {
			continue
		}
		for j := 0; j < count; j++ {
			fmt.Fprintln(writer, i+1)
		}
	}
}
