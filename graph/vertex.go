package graph

type Vertex interface {
	GetBody()		string
	GetTimestamp()	uint64
}

type vertex struct {
	body			string
	timestamp		uint64
}

func (v *vertex) GetBody() string {
	return v.body
}

func (v *vertex) GetTimestamp() uint64 {
	return v.timestamp
}

func createVertex(theBody string, ts uint64) Vertex {
	return &vertex {
		body: theBody,
		timestamp: ts,
	}
}
