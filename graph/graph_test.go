package graph

import (
	"testing"
)

const sqlQuery = "SELECT * FROM ??"
const mongoQuery = "db.collection.find({id: ??})"
const graphQLquery = `query {
  repository(owner: "graphql", name: "graphql-js"){
    name
    description
  }
}`

const timestampForAnyRecord = 555

func getRecord(body string, ts uint64) Record {
	return Record{
		Tid: 1,
		Body: body,
		Timestamp: ts,
		Params: make(map[string]interface{}),
	}
}

func TestGraphInitialization(t *testing.T) {
	t.Log("INIT: start testing graph without nodes")
	theHash := getHash(sqlQuery)
	record := getRecord(sqlQuery, timestampForAnyRecord)
	//graph := createGraph(sqlQuery, theHash)
	graph := newGraph(record)
	if graph.current != theHash || graph.root != theHash {
		t.Error("root and current node should be equal to last node hash")
	}
	t.Log("current vertex and root vertex pointer are correct")
	if len(graph.vertices) != 1 {
		t.Error("length of verteces map is not equal to 1")
	}
	if graph.targetsOfVertex[theHash].count() != 0 {
		t.Error("graph with single vertex has phantom edges")
	}
	t.Log("graph has single vertex without edges")
	if graph.finished {
		t.Error("graph is not finished but has broken flag")
	}
}

func TestGraphWithTwoNodes(t *testing.T) {
	t.Log("INIT: Start testing two nodes graph")
	rootHash := getHash(sqlQuery)
	hashOfSecondNode := getHash(mongoQuery)
	graph := newGraph(getRecord(sqlQuery, timestampForAnyRecord))
	graph.RegisterRecord(getRecord(mongoQuery, timestampForAnyRecord + 15))

	if graph.finished || graph.root != rootHash || graph.current != hashOfSecondNode {
		t.Error("smth went wrong!")
	}
	if len(graph.vertices) != 2 {
		t.Error("more than 2 nodes in this case")
	}

	if graph.targetsOfVertex[rootHash].count() != 1 && graph.targetsOfVertex[hashOfSecondNode].count() != 0 {
		t.Error("error weights and edges!")
	}
}