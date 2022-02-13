package main

import "fmt"

type laptopSize float64

func (this laptopSize) getSizeOfLaptop() laptopSize {
	return this
}

func main() {
	var lp laptopSize = 3.5

	fmt.Println(lp.getSizeOfLaptop())
}
