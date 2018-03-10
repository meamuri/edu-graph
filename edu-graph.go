package main

import (
	"fmt"
	"github.com/meamuri/edu-graph/graph"
)

func main() {
	r := graph.Record{Tid: 100, Body: "SELECT ?? FROM table_name", Timestamp: 12, Params: make(map[string]interface{})}
	g := graph.NewGraph(r)
	if g == nil {
		fmt.Printf("error")
	}
	fmt.Printf("hi")
}
