package graph

type Edges interface {
	computeEdge(previousTs, ts, target uint64)
}

type edges struct {
	weights map[uint64]Weight
}

// method should be invoked with vertex initialization
func createEdges() *edges {
	return &edges {
		weights: make(map[uint64]Weight),
	}
}

// edge will be created if does not exist
// otherwise edge weight will be recomputed
func (e *edges) computeEdge(previousTs, ts, target uint64) {
	if w, ok := e.weights[target]; ok {
		w.Recompute(previousTs, ts)
	} else {
		e.weights[target] = createWeight(previousTs, ts)
	}
}