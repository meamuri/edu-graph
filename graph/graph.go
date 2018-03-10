package graph

type Graph interface {
	RegisterRecord(record Record)	bool
	String() string
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
		root:            theHash,								// first elem is root of graph
		current:         theHash,								// also first elem is current elem
		finished:        false,									// true after special signal
		vertices:        make(map[uint64]Vertex),
		sourcesOfVertex: make(map[uint64]map[uint64]weight),
		targetsOfVertex: make(map[uint64]map[uint64]weight),
	}
	res.vertices[theHash] = createVertex(body)
	// empty map means, that RootVertex (theHash) does not contain edges yet
	// maps initialized, but should be empty now (we have single root
	res.sourcesOfVertex[theHash] = make(map[uint64]weight)
	res.targetsOfVertex[theHash] = make(map[uint64]weight)
	return &res
}

func (g *graph) String() string {
	return "graph"
}

func (g *graph) RegisterRecord(record Record) bool {
	sum := getHash(record.Body)
	//if _, ok := g.vertices[sum]; !ok {
	//	return false
	//}

	if _, ok := g.targetsOfVertex[g.current][sum]; !ok {

	}

	// 1.


	return true
}

func NewGraph(record Record) Graph {
	sum := getHash(record.Body)
	return newGraph(record.Body, sum)
}