package avito

type Segment struct {
	SegmentName string `json:"segment_name" db:"segment_name" binding:"required"`
}
