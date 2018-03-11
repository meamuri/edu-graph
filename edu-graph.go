package main

import (
	"fmt"
)

func main() {
	r := Record{Tid: 100, Body: "SELECT ?? FROM table_name", Timestamp: 12, Params: make(map[string]interface{})}
	g := NewGraph(r)
	if g == nil {
		fmt.Printf("error")
	}
	fmt.Printf("hi")
}
