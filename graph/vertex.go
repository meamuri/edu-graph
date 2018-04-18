package graph

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

func (v *vertex) RegisterNewTimestamp(timestamp uint64) {
	v.timestamp = timestamp
}

func createVertex(theBody string, ts uint64) *vertex {
	return &vertex {
		body: theBody,
		timestamp: ts,
	}
}
