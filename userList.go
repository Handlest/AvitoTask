package avito

type UserList struct {
	UserId             int      `json:"userId" db:"userId" binding:"required"`
	SegmentNamesAdd    []string `json:"segment_names_add" db:"segment_names_add" binding:"required"`
	SegmentNamesRemove []string `json:"segment_names_remove" db:"segment_names_remove" binding:"required"`
	Added              string   `json:"added" db:"added"`
	Expiry             string   `json:"expiry" db:"expiry"`
}
