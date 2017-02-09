package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a height for the staircase: ")
	fmt.Scanf("%d", &n)

	for i := n; i > 0; i-- {
		for j := 1; j <= n; j++ {
			if j < i {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
}
