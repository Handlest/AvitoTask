package repository

import (
	avito "AvitoTask"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user avito.User) error
	GetOperations(user avito.UserInfo) ([]avito.Operation, error)
	DeleteUser(userId int, segmentName string) error
}

type Segment interface {
	CreateSegment(segment avito.Segment) (bool, error)
	GetSegments(userId int) ([]avito.Segment, error)
	DeleteSegment(segmentName string) error
}
type Operation interface {
}
type Repository struct {
	User
	Segment
	Operation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{User: NewCreateUserPostgres(db), Segment: NewSegmentPostgres(db)}
}
