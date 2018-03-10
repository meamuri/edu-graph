package graph

type Graph interface {
	RegisterRecord(record Record)	bool
}

type graph struct {
	root     		uint64 // hash of root vertex
	current  		uint64 // hash of current vertex
	finished 		bool
	vertices		map[uint64]Vertex
	sourcesOfVertex	map[uint64]map[uint64]weight
	targetsOfVertex	map[uint64]map[uint64]weight
}

func newGraph(body string, theHash uint64) *graph {
	res := graph{
		root:            theHash,
		current:         theHash,
		finished:        false,
		vertices:        make(map[uint64]Vertex),
		sourcesOfVertex: make(map[uint64]map[uint64]weight),
		targetsOfVertex: make(map[uint64]map[uint64]weight),
	}
	res.vertices[theHash] = &vertex{}
	return &res
}

func (g *graph) RegisterRecord(record Record) bool {

	return true
	//sum := getHash(record.GetBody())
}

func NewGraph(record Record) Graph {
	sum := getHash(record.GetBody())
	return newGraph(record.GetBody(), sum)
}