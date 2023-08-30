package avito

type Operation struct {
	UserId         int    `json:"userId"`
	SegmentName    string `json:"segment_name"`
	Operation_type string `json:"operation_type"`
	DateTime       string `json:"dateTime"`
}
