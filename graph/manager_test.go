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
	t.Log("INIT: manager empty, try register first Tid.")
	m := CreateManager()
	m.ManageRecord(constRecord)
	if false {
		t.Error("something wrong!")
	}
}