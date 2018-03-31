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

func TestManagerEmpty(t *testing.T) {
	t.Log("INIT: manager empty, check fields.")
	m := createManager()
	if m.currentElement != -1 || len(m.graphToTheirTIDs) != 0 {
		t.Error("managare has incorrect current element")
	}
}

func TestManagerWithSingleRecord(t *testing.T) {
	t.Log("INIT: manager empty, try register first Tid.")
	m := createManager()
	m.ManageRecord(constRecord)
	if ! func() bool {
		return m.currentElement == 0 &&
			len(m.graphToTheirTIDs) == 1 &&
			m.graphToTheirTIDs[constRecord.Tid] == 0
	}() {
		t.Error("managare has incorrect current element")
	}
}