package main

import "time"

func StartTicking(shouldBeSnapshoted chan<- bool, periodInSeconds int) {
	t := time.NewTicker(time.Second * time.Duration(periodInSeconds))
	for range t.C {
		shouldBeSnapshoted <- true
	}
}
