package graph

type Record struct {
	Tid int `json:"tid"`
	Body string `json:"body"`
	Timestamp int64 `json:"timestamp"`
	Params map[string]interface{} `json:"params"`
}

type Weight struct {

}

type Graph struct {
	Root int
	Vertexes map[int]map[int]Weight
	Current int
}
