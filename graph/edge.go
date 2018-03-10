package graph

type Edge interface {
	Source()		Vertex
	Target()		Vertex
	GetWeight()		Weight
}

type edge struct {
	source	Vertex
	target	Vertex
	wgt		Weight
}

func (e *edge) Source() Vertex {
	return e.source
}
