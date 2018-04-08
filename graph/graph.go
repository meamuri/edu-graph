package graph

// data types {


// Graph (discrete mathematics), a set of vertices and edges
type Graph interface {
	RegisterRecord(record Record)	bool
	String() string
}

type graph struct {
	root     		uint64 // hash of root vertex
	current  		uint64 // hash of current vertex
	finished 		bool
	vertices		map[uint64]Vertex
	targetsOfVertex	map[uint64]Edges
}

// } .. dataTypes

// utils {

// add vertex into graph and initialize map of targets for new vertex
// if vertex exists, do nothing.
// return `true`, if vertex was added, `false` otherwise
func (g *graph) addVertex(body string, ts uint64, address uint64) bool {
	if _, ok := g.vertices[address]; ok {
		return false // vertex exists, return and notify that we did not add vertex
	}
	g.vertices[address] = createVertex(body, ts)
	g.targetsOfVertex[address] = createEdges()
	return true
}


func createGraph(body string, ts uint64, theHash uint64) *graph {
	res := graph{
		root:            theHash,								// first elem is root of graph
		current:         theHash,								// also first elem is current elem
		finished:        false,									// true after special signal
		vertices:        make(map[uint64]Vertex),
		targetsOfVertex: make(map[uint64]Edges),
	}
	// code blow is very similar addVertex function,
	// but we sure that initial vertex does not exist
	// and we can skip `if` operator [optimization, ha-ha]
	res.vertices[theHash] = createVertex(body, ts)
	// empty map means, that RootVertex (theHash) does not contain edges yet
	// maps initialized, but should be empty now (we have single root
	res.targetsOfVertex[theHash] = createEdges()
	return &res
}

func (g *graph) addEdge(to uint64) {
	if to == g.current {
		panic("target vertex and current vertex have same values!")
	}
	from := g.current
	prevTs := g.vertices[from].GetTimestamp()
	ts := g.vertices[to].GetTimestamp()

	g.targetsOfVertex[from].computeEdge(prevTs, ts, to)
}

// } .. utils

// interface implementation {

func (g *graph) String() string {
	return "graph"	// TODO: impl string representation
}

func (g *graph) RegisterRecord(record Record) bool {
	sum := getHash(record.Body)
	g.addVertex(record.Body, record.Timestamp, sum) // we can ignore `bool` result
	g.addEdge(sum)
	g.current = sum // now we stay here
	return true
}

// } .. interface implementation

// Api {

// this factory will be used in tests,
// for access to fields and assertion
func newGraph(record Record) *graph {
	sum := getHash(record.Body)
	return createGraph(record.Body, record.Timestamp, sum)
}
func NewGraph(record Record) Graph {
	return newGraph(record)
}

// } .. api