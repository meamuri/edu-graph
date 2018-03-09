package api

type Record interface {
	GetTID() 		int
	GetBody() 		string
	GetTimestamp() 	uint64
	GetParams() 	map[string]interface{}
}

type record struct {
	tid			int						`json:"tid"`
	body 		string					`json:"body"`
	timestamp 	uint64					`json:"timestamp"`
	params 		map[string]interface{} 	`json:"params"`
}