package main

import (
	"fmt"
	"net"
	"os"
	"encoding/json"
	"github.com/meamuri/edu-graph/graph"
)

const (
	HOST = "localhost"
	PORT = "514"
	TYPE = "tcp"
)

func StartListening(c chan<- graph.Record, b chan<-bool) {
	hostPlusPort := HOST + ":" + PORT
	l, err := net.Listen(TYPE, hostPlusPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	defer func() {
		b<-true
	}()
	fmt.Println("Listening on " + hostPlusPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn, c)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, c chan<- graph.Record) {
	d := json.NewDecoder(conn)

	var msg graph.Record
	err := d.Decode(&msg)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	c <- msg
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
