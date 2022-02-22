package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://yahoo.com",
	}

	c := make(chan string)

	for _, v := range urls {
		go callUrl(v, c) // creates a new goroutine
	}

	/***
	 * Make persistant connection with channel and the program won't stop automatically.
	 */
	// for v := range c {
	// 	fmt.Println(v)
	// }

	// This loop breaks the channel connection since the loop will exhaust.
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c) // Blocking code
	}
}

func callUrl(u string, c chan string) {
	fmt.Println("calling", u)
	_, err := http.Get(u) // blocking code

	if err != nil {
		fmt.Println("Error:", err)
	}

	c <- "Go it"
}
