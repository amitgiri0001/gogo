package main

import (
	"fmt"
	"io"
	"net/http"
)

//Creating writer function for react the
type logger struct{}

func (logger) Write(b []byte) (n int, err error) {
	fmt.Print(string(b))

	return len(b), nil
}

func main() {
	data, _ := http.Get("https://api.mfapi.in/mf/122639")
	lg := logger{}
	io.Copy(lg, data.Body)
}

// Reading data body from an HTTP call
func main_1() {
	data, _ := http.Get("https://api.mfapi.in/mf/122639")

	d, _ := io.ReadAll(data.Body)

	fmt.Println(string(d))
}
