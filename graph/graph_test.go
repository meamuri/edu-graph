package graph

import (
	"testing"
)

const someSQL = "SELECT * FROM ??"
const someQuery  = "db.collection.find({id: ??})"
const someGraphQL = ""

func TestGraphInitialization(t *testing.T) {
	t.Log("graph is empty. Try init with first record")
	theHash := getHash(someSQL)
	graph := createGraph(someSQL, theHash)
	if graph.current != theHash || graph.root != theHash {
		t.Error("root and current node should be equal to last node hash")
	}
	t.Log("current vertex and root vertex pointer are correct")
	if len(graph.vertices) != 1 {
		t.Error("length of verteces map is not equal to 1")
	}
	if len(graph.targetsOfVertex[theHash]) != 0 {
		t.Error("graph with single vertex has phantom edges")
	}
	t.Log("graph has single vertex without edges")
	if graph.finished {
		t.Log("graph is not finished but has broken flag")
	}
}

func TestGraphWithTwoNodes(t *testing.T) {
	t.Log("graph is empty. Try init with first record")
	theHash := getHash(someSQL)
	hashOfSecondNode := getHash(someQuery)
	graph := createGraph(someSQL, theHash)
	graph.addVertex(someQuery, hashOfSecondNode)
}