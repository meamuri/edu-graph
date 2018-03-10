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
