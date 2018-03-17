package main

import (
	"fmt"
)

// rc is a Record Chan
// sc is a Snapshot Chan
// fc is a Finish Chan
func controlFlow(m Manager, rc chan Record, sc, fc chan bool) {
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
	fmt.Printf("Start execution")

	m := CreateManager()
	rc := make(chan Record)
	cs := make(chan bool)
	fc := make(chan bool)

	// we should stop execution before end of program
	defer func(){
		fc <- true
	}()

	go controlFlow(m, rc, cs, fc)
	// crete sysLog
	// create neo4j

}
