package avito

type User struct {
	UserId  int    `json:"userId" binding:"required"`
	Segment string `json:"segment" binding:"required"`
	Added   string `json:"added"`
	Expiry  string `json:"expiry"`
}
