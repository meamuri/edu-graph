package graph

type Weight struct {
	Min int
	Max int
	Delta int
	Dispersion int
}

type Vertex struct {
	Children map[int]Weight
}

type Graph struct {
	Root int	// hash of root vertex
	Current int	// hash of current vertex
	Vertexes map[int]Vertex
	Finished bool
}
