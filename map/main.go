package main

import "fmt"

func main() {
	//var stocks map[string]int -> Only declaration
	// stocks := make(map[string]int) // declaration + initialization
	stocks := map[string]int{ // declaration + initialization + definition
		"pltr": 100,
	}
	stocks["goog"] = 200
	stocks["fb"] = 500

	for k, v := range stocks {
		fmt.Println(k, " value is", v)
	}

}
