package graph

// data types {


// Graph (discrete mathematics), a set of vertices and edges
type Graph interface {
	RegisterRecord(record Record)
	String() string
}

type graph struct {
	root     		uint64 // hash of root vertex
	current  		uint64 // hash of current vertex
	finished 		bool
	vertices		map[uint64]*vertex
	targetsOfVertex	map[uint64]*edges
}

// } .. dataTypes

// utils {

func (g *graph) getEdgesOf(vertexHash uint64) *edges {
	return g.targetsOfVertex[vertexHash]
}

func (g *graph) getTsOf(vertexHash uint64) uint64 {
	return g.vertices[vertexHash].GetTimestamp()
}

func (g *graph) getEdgesOfCurrent() *edges {
	return g.targetsOfVertex[g.current]
}

func (g *graph) getTimestampOfCurrent() uint64 {
	return g.vertices[g.current].GetTimestamp()
}

// add vertex into graph and initialize map of targets for new vertex
// if vertex exists, do nothing.
// return `true`, if vertex was added, `false` otherwise
func (g *graph) addVertex(body string, timestamp uint64, vertexHash uint64) bool {
	if _, ok := g.vertices[vertexHash]; ok {
		return false // vertex exists, return and notify that we did not add vertex
	}
	g.vertices[vertexHash] = createVertex(body, timestamp)
	g.targetsOfVertex[vertexHash] = createEdges()
	return true
}


func createGraph(body string, ts uint64, theHash uint64) *graph {
	res := graph{
		root:            theHash,								// first elem is root of graph
		current:         theHash,								// also first elem is current elem
		finished:        false,									// true after special signal
		vertices:        make(map[uint64]*vertex),
		targetsOfVertex: make(map[uint64]*edges),
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

// Create edge between current and next vertices. Panic, if current and target is the same vertices.
func (g *graph) addEdge(sourceVertex, targetVertex uint64) {
	timestamp := g.vertices[targetVertex].GetTimestamp()
	if sourceVertex == targetVertex {
		g.addSelfLoop(sourceVertex, timestamp)
		return
	}
	previousTimestamp := g.vertices[sourceVertex].GetTimestamp()
	g.targetsOfVertex[sourceVertex].computeEdge(previousTimestamp, timestamp, targetVertex)
}

// Create loop edge at last vertex
func (g *graph) addSelfLoop(vertexHash, timestamp uint64) {
	previousTimestamp := g.getTsOf(vertexHash)
	g.getEdgesOf(vertexHash).computeEdge(previousTimestamp, timestamp, vertexHash)
}

// } .. utils

// interface implementation {

func (g *graph) String() string {
	return "graph"	// TODO: impl string representation
}

func (g *graph) RegisterRecord(record Record) {
	vertexHash := getHash(record.Body)
	// if hashes are equals, create loop edge and recompute last vertex timestamp
	g.addVertex(record.Body, record.Timestamp, vertexHash) // we can ignore `bool` result
	g.addEdge(g.current, vertexHash)
	g.vertices[vertexHash].RegisterNewTimestamp(record.Timestamp)
	g.current = vertexHash // now we stay here
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