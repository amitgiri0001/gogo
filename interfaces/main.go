package main

import "fmt"

type Animal interface {
	speak() string
}

type Cat struct{}
type Dog struct{}

func (Dog) speak() string {
	return "Barks!"
}

func (Cat) speak() string {
	return "Meow!"
}

func print(a Animal) {
	fmt.Println(a.speak())
}

func main() {
	c := Cat{}
	print(c)

	d := Dog{}
	print(d)
}
