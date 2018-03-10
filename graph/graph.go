package graph

type Graph struct {
	Root int	// hash of root vertex
	Current int	// hash of current vertex
	Vertexes map[int]Vertex
	Finished bool
}
