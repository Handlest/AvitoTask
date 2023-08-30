package avito

type Operation struct {
	UserId   int    `json:"userId"`
	Segment  string `json:"segment"`
	Type     string `json:"type"`
	DateTime string `json:"dateTime"`
}
