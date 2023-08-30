package avito

type UserInfo struct {
	UserId    int    `json:"userId" db:"userId" binding:"required"`
	StartDate string `json:"start" db:"start" binding:"required"`
	EndDate   string `json:"end" db:"end" binding:"required"`
}
