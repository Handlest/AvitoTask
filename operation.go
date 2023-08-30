package avito

type Operation struct {
	UserId        int    `json:"userId"  db:"user_id"`
	SegmentName   string `json:"segment_name"  db:"segment_name"`
	OperationType string `json:"operation_type" db:"operation_type"`
	DateTime      string `json:"dateTime" db:"operation_date"`
}
