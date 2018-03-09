package api

type Record interface {
	GetTID() 		int
	GetBody() 		string
	GetTimestamp() 	uint64
	GetParams() 	map[string]interface{}
}

type record struct {
	Tid       int                    `json:"tid"`
	Body      string                 `json:"body"`
	Timestamp uint64                 `json:"timestamp"`
	Params    map[string]interface{} `json:"params"`
}

func (r *record) GetTid() int {
	return r.Tid
}

func (r *record) GetBody() string {
	return r.Body
}

func (r *record) GetTimestamp() uint64 {
	return r.Timestamp
}

func (r *record) GetParams() map[string]interface{} {
	return r.Params
}
