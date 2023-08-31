package avito

type Segment struct {
	SegmentName string `json:"segment_name" db:"segment_name" binding:"required"`
}

type SegmentInSwagger struct {
	SegmentName string `json:"segmentName" db:"segmentName" binding:"required"`
}
