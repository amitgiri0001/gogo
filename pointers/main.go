package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	newPerson := Person{firstName: "Amit_1", lastName: "Giri_1"}
	fmt.Printf("Fresh assignment %+v \n", newPerson)

	newPerson.changeName("New Amit", "New Giri")
	fmt.Printf("After changeName %+v \n", newPerson)

	passByValue(newPerson)
	fmt.Printf("After passByValue  %+v \n", newPerson)

	personPointer := &newPerson
	passByRef(personPointer)
	fmt.Printf("After passByRef  %+v \n", newPerson)

}

func passByRef(personPointer *Person) {
	// One way to update value of a pointer
	personPointer.firstName = "Rony"
	//(*personPointer).firstName = "Ron"
	// Another way to update value of a pointer
	(*personPointer).lastName = "hut"

	fmt.Printf("In passByRef - by value %+v \n", *personPointer)
	fmt.Printf("In passByRef - by address %+v \n", personPointer)
}

func passByValue(newPerson Person) {
	newPerson.firstName = "New name"
	fmt.Printf("In passByValue %+v \n", newPerson)
}

func (p Person) changeName(firstName string, lastName string) {
	p.firstName = firstName
	p.lastName = lastName
	fmt.Printf("In changeName %+v \n", p)
}
