package graph



type Record struct {
	Tid int `json:"tid"`
	Body string `json:"body"`
	Timestamp int64 `json:"timestamp"`
	Params map[string]interface{} `json:"params"`
}

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
