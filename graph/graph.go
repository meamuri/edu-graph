package graph

import (
	"hash/fnv"
)

type Graph interface {

}

type graph struct {
	Root uint64				// hash of root vertex
	Current uint64			// hash of current vertex
	Vertexes map[uint64]Vertex
	Finished bool
}

func newGraph(body string) *graph {
	h := fnv.New64()
	h.Write([]byte(body))
	sum := h.Sum64()
	return &graph{
		Root: sum,
		Current: sum,
		Finished: false,
		Vertexes: make(map[uint64]Vertex),
	}
}

func NewGraph(body string) Graph {
	return newGraph(body)
}