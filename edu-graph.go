package main

import (
	"fmt"
	"github.com/meamuri/edu-graph/graph"
	"time"
)

const SnapshotPeriod = 15

// rc is a Record Chan
// sc is a Snapshot Chan
// fc is a Finish Chan
func controlFlow(manager graph.Manager, record <-chan graph.Record, snapshot, finish <-chan bool) {
	for {
		select {
		case r := <-record:
			manager.ManageRecord(r)
		case <-snapshot:
			fmt.Println("Should be snapshoted now: ", time.Now())
			// snapshot
		case <-finish:
			// snapshot
			return
		}
	}
}

func main() {
	graphManager := graph.CreateManager()
	recordsChannel := make(chan graph.Record)
	snapShotSignalChannel := make(chan bool)
	finishSignalChannel := make(chan bool)

	// we should stop execution before end of program
	defer func() {
		finishSignalChannel <- true
	}()

	go controlFlow(graphManager, recordsChannel, snapShotSignalChannel, finishSignalChannel)

	listenerErrorSignal := make(chan bool)
	receivedRecords := make (chan graph.Record)
	go StartListening(receivedRecords, listenerErrorSignal)
	go StartTicking(snapShotSignalChannel, SnapshotPeriod)

	for {
		select {
			case record := <-receivedRecords:
				fmt.Println(record)
				recordsChannel <- record // pipe message to manager
			case <-listenerErrorSignal:
				fmt.Printf("good bye")
				break
		}
	}
	// crete sysLog
	// create neo4j

}
