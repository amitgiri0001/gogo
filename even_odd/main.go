package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for _, n := range numbers {
		nType := "odd"

		if n%2 == 0 {
			nType = "even"
		}

		fmt.Println(n, "is", nType)
	}
}
