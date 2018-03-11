package graph

import (
	"testing"
)

var constRecord = Record {
	Tid: 1,
	Body: "SELECT ??",
	Timestamp: 10000,
	Params: map[string]interface{}{
		"select": "*",
	},
}

func TestManagerRegisterFirstTID(t *testing.T) {
	t.Log("manager empty, try register first Tid.")
	m := CreateManager()
	m.ManageRecord(constRecord)
}