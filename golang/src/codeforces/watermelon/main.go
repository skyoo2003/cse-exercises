package main

import "fmt"

func main() {
	var w int
	fmt.Scanf("%d", &w)

	if w == 1 {
		fmt.Println("NO")
		return
	}
	for i := 2; i <= w/2; i++ {
		if w%i == 0 && w%2 == 0 {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
