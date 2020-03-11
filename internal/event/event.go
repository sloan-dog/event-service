package event

type Event struct {
	Id      string                 `json:"id"`
	Name    string                 `json:"name"`
	Context []string               `json:"context"`
	Data    map[string]interface{} `json:"data"`
	Time    int64                  `json:"time"` // big ole unix timestamp
}
