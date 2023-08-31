package avito

type User struct {
	UserId      int    `json:"userId" db:"userId" binding:"required"`
	SegmentName string `json:"segment_name" db:"segment_name" binding:"required"`
	Added       string `json:"added" db:"added"`
	Expiry      string `json:"expiry" db:"expiry"`
}

type UserInfo struct {
	UserId    int    `json:"userId" db:"userId" binding:"required"`
	StartDate string `json:"start" db:"start" binding:"required"`
	EndDate   string `json:"end" db:"end" binding:"required"`
}

type UserList struct {
	UserId             int      `json:"userId" db:"userId" binding:"required"`
	SegmentNamesAdd    []string `json:"segment_names_add" db:"segment_names_add" binding:"required"`
	SegmentNamesRemove []string `json:"segment_names_remove" db:"segment_names_remove" binding:"required"`
	Added              string   `json:"added" db:"added"`
	Expiry             string   `json:"expiry" db:"expiry"`
}

type UserSegmentWithId struct {
	UserId      int    `json:"userId" db:"userId" binding:"required"`
	SegmentName string `json:"segmentName" db:"segmentName" binding:"required"`
}

type UserIdOnly struct {
	UserId int `json:"userId" db:"userId" binding:"required"`
}
