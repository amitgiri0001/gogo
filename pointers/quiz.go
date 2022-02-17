// package main

// import "fmt"

// func main() {
// 	name := "Bill"

// 	fmt.Println(*&name)
// }

package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
