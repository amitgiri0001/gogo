package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer getTime(1) // it will execute after the function return at last
	defer getTime(2) // it will execute after the function return at second last
	log.Println("from main()", time.Now())
}

func getTime(i int) {
	defer fmt.Println(i, " deferred", time.Now())
	log.Println(1, " form getTime()", time.Now())
}
