package avito

type User struct {
	UserId      int    `json:"userId" binding:"required"`
	SegmentName string `json:"segment_name" binding:"required"`
	Added       string `json:"added"`
	Expiry      string `json:"expiry"`
}
