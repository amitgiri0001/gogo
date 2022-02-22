package main

import (
	"fmt"
	"io"
	"os"
)

type logger struct{}

func (logger) Write(b []byte) (n int, err error) {
	fmt.Println(string(b))
	return len(b), nil
}

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lg := logger{}
	io.Copy(lg, file)
}
