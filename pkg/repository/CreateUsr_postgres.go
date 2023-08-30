package repository

import (
	avito "AvitoTask"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type CreateUserPostgres struct {
	db *sqlx.DB
}

func NewCreateUserPostgres(db *sqlx.DB) *CreateUserPostgres {
	return &CreateUserPostgres{db: db}
}

func (r *CreateUserPostgres) CreateUser(user avito.User) (int, error) {
	fmt.Println("CreateUserFromRepository!!")

	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_Id, segment_name, added, expiry) VALUES ($1, $2, $3, $4) RETURNING user_id", usersTable)
	row := r.db.QueryRow(query, user.UserId, user.Segment, time.Now(), user.Expiry)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
