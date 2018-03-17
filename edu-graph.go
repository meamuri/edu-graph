package main

import (
	"fmt"
	"github.com/meamuri/edu-graph/graph"
	"github.com/meamuri/edu-graph/connectivity"
)

// rc is a Record Chan
// sc is a Snapshot Chan
// fc is a Finish Chan
func controlFlow(m graph.Manager, rc <-chan graph.Record, sc, fc <-chan bool) {
	for {
		select {
		case r := <- rc:
			m.ManageRecord(r)
		case <- sc:
			// snapshot
		case <- fc:
			// snapshot
			return
		}
	}
}

func main() {
	//r := Record{Tid: 100, Body: "SELECT ?? FROM table_name", Timestamp: 12, Params: make(map[string]interface{})}
	//g := NewGraph(r)
	//if g == nil {
	//	fmt.Printf("error")
	//}
	//fmt.Printf("hi")
	fmt.Printf("Start execution\n")

	m := graph.CreateManager()
	rc := make(chan graph.Record)
	cs := make(chan bool)
	fc := make(chan bool)

	// we should stop execution before end of program
	defer func(){
		fc <- true
	}()

	go controlFlow(m, rc, cs, fc)

	bc := make(chan bool)
	stringChan := make (chan string)
	go connectivity.StartListening(stringChan, bc)

	for {
		select {
			case s := <-stringChan:
				fmt.Printf(s)
			case <-bc :
				fmt.Printf("good bye")
				break
		}
	}
	// crete sysLog
	// create neo4j

}
