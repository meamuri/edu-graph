package graph

type Vertex interface {
	GetHash()		uint64
}

type vertex struct {
	hash			uint64
}

func (v *vertex) GetHash() uint64 {
	return v.hash
}

type Weight interface {
	GetMin()		int
	GetMax()		int
	GetDelta()		int
	GetDispersion()	int
}

type weight struct {
	min			int
	max			int
	delta		int
	dispersion 	int
}

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


//type Vertex struct {
//	Children map[int]Weight
//}

type Graph struct {
	Root int	// hash of root vertex
	Current int	// hash of current vertex
	Vertexes map[int]Vertex
	Finished bool
}
