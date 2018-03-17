package graph

type Vertex interface {
	GetBody()		string
}

type vertex struct {
	body			string
}

func (v *vertex) GetBody() string {
	return v.body
}

func createVertex(theBody string) Vertex {
	return &vertex {
		body: theBody,
	}
}