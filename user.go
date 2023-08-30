package avito

type User struct {
	UserId      int    `json:"userId" db:"userId" binding:"required"`
	SegmentName string `json:"segment_name" db:"segment_name" binding:"required"`
	Added       string `json:"added" db:"added"`
	Expiry      string `json:"expiry" db:"expiry"`
}
