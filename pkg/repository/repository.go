package repository

import (
	avito "AvitoTask"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user avito.User) (int, error)
}

type Segment interface {
}
type Operation interface {
}
type Repository struct {
	User
	Segment
	Operation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{User: NewCreateUserPostgres(db)}
}
