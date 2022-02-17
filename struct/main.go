package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	newPerson := Person{"Amit", "Giri"}
	fmt.Printf("%+v \n", newPerson)

	newPerson = Person{firstName: "Amit_1", lastName: "Giri_1"}
	fmt.Printf("%+v \n", newPerson)

	newPerson.firstName = "Amit_2"
	newPerson.lastName = "Giri_2"
	fmt.Printf("%+v \n", newPerson)

}
