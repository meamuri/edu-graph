package graph

import (
	"fmt"
)

// single record of logical business transaction
type Record struct {
	// Transaction id
	Tid       int                    `json:"tid"`

	// Query body
	Body      string                 `json:"body"`

	// Timestamp of execution
	Timestamp uint64                 `json:"timestamp"`

	// Parameters of execution
	Params    map[string]interface{} `json:"params"`
}

// String representation of record
func (r *Record) String() string {
	return fmt.Sprintf("%v %v %v", r.Tid, r.Body, r.Timestamp)
}
