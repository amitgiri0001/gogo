package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	/**
	 * Sequential assignment
	 */
	newPerson := Person{"Amit", "Giri"}
	fmt.Printf("%+v \n", newPerson)

	/**
	 * JS Object like assignment
	 */
	newPerson = Person{firstName: "Amit_1", lastName: "Giri_1"}
	fmt.Printf("%+v \n", newPerson)

	/**
	 * instance values
	 */
	newPerson.firstName = "Amit_2"
	newPerson.lastName = "Giri_2"
	fmt.Printf("%+v \n", newPerson)

	/**
	 * Proves that GO is a pass by "value" language not by "ref"
	 */
	person2 := newPerson
	fmt.Printf("%+v \n", person2)
	person2.firstName = "Not Amit"
	fmt.Printf("%+v \n", person2)
	fmt.Printf("%+v \n", newPerson)

}
