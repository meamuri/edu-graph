package main

// syslog record, single action of logical business transaction
type Record struct {
	// transaction id
	Tid       int                    `json:"tid"`

	// query body -- base unit of transaction
	Body      string                 `json:"body"`

	// timestamp of action
	Timestamp uint64                 `json:"timestamp"`

	// params of action
	Params    map[string]interface{} `json:"params"`
}
