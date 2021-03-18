// nolint
package binomialcoefficient

import (
	"fmt"
	"math"
	"os"
)

func binRecursive(n, k int) int {
	if k == 0 || n == k {
		return 1
	}
	return binRecursive(n-1, k-1) + binRecursive(n-1, k)
}

// BinRecursive My solution
func BinRecursive(in *os.File) {
	var n, k int
	fmt.Fscanf(in, "%d %d", &n, &k)
	fmt.Println(binRecursive(n, k))
}

func binDP(n, k int) int {
	table := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		table[i] = make([]int, k+1)
	}

	for i := 0; i <= n; i++ {
		lm := int(math.Min(float64(i), float64(k)))
		for j := 0; j <= lm; j++ {
			if j == 0 || j == i {
				table[i][j] = 1
			} else {
				table[i][j] = table[i-1][j-1] + table[i-1][j]
			}
		}
	}
	return table[n][k]
}

// BinDP My solution
func BinDP(in *os.File) {
	var n, k int
	fmt.Fscanf(in, "%d %d", &n, &k)
	fmt.Println(binDP(n, k))
}
