package avito

type Operation struct {
	UserId      int    `json:"userId"`
	SegmentName string `json:"segment_name"`
	Type        string `json:"type"`
	DateTime    string `json:"dateTime"`
}
