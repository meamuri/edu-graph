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
	targetsOfVertex	map[uint64]map[uint64]Weight
}

// } .. dataTypes

// utils {

// add vertex into graph and initialize map of targets for new vertex
// if vertex exists, do nothing.
// return `true`, if vertex was added, `false` otherwise
func (g *graph) addVertex(body string, address uint64) bool {
	if _, ok := g.vertices[address]; ok {
		return false // vertex exists, return and notify that we did not add vertex
	}
	g.vertices[address] = createVertex(body)
	g.targetsOfVertex[address] = make(map[uint64]Weight)
	return true
}

// add edge between two vertices: i->j
// edge weight will computed in according ``ts`` param [timestamp]
// if edge exists, recompute weight
func (g *graph) computeEdgeWeight(i, j, ts uint64) {
	if val, ok := g.targetsOfVertex[i][j]; ok {
		val.Recompute(ts)
	} else {
		g.targetsOfVertex[i][j] = CreateWeight(ts)
	}
}

func newGraph(body string, theHash uint64) *graph {
	res := graph{
		root:            theHash,								// first elem is root of graph
		current:         theHash,								// also first elem is current elem
		finished:        false,									// true after special signal
		vertices:        make(map[uint64]Vertex),
		targetsOfVertex: make(map[uint64]map[uint64]Weight),
	}
	// code blow is very similar addVertex function,
	// but we sure that initial vertex does not exist
	// and we can skip `if` operator [optimization, ha-ha]
	res.vertices[theHash] = createVertex(body)
	// empty map means, that RootVertex (theHash) does not contain edges yet
	// maps initialized, but should be empty now (we have single root
	res.targetsOfVertex[theHash] = make(map[uint64]Weight)
	return &res
}

// } .. utils

// interface implementation {

func (g *graph) String() string {
	return "graph"	// TODO: impl string representation
}

func (g *graph) RegisterRecord(record Record) bool {
	sum := getHash(record.Body)
	g.addVertex(record.Body, sum) // we can ignore `bool` result
	g.computeEdgeWeight(g.current, sum, record.Timestamp)
	g.current = sum // now we stay here
	return true
}

// } .. interface implementation

// Api {

func NewGraph(record Record) Graph {
	sum := getHash(record.Body)
	return newGraph(record.Body, sum)
}

// } .. api