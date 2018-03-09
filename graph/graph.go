package graph

// Vertex part

type Vertex interface {
	GetHash()		uint64
}

type vertex struct {
	hash			uint64
}

func (v *vertex) GetHash() uint64 {
	return v.hash
}

// End vertex part

// Weight part

type Weight interface {
	GetMin()		int
	GetMax()		int
	GetDelta()		int
	GetDispersion()	int

	Recompute(ts uint64)
}

type weight struct {
	min			int
	max			int
	delta		int
	dispersion 	int
}

func (w *weight) GetMin() int {
	return w.min
}
func (w *weight) GetMax() int {
	return w.max
}
func (w *weight) GetDelta() int {
	return w.delta
}
func (w *weight) GetDispersion() int {
	return w.dispersion
}
func (w *weight) Recompute(ts uint64) {

}

// End Weight part

// Edge part

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

// end edge part

type Graph struct {
	Root int	// hash of root vertex
	Current int	// hash of current vertex
	Vertexes map[int]Vertex
	Finished bool
}
