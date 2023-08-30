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

func (r *CreateUserPostgres) CreateUser(user avito.User) error {
	var expiryDate string
	var additionTime = time.Now()
	if user.Expiry == "" {
		expiryDate = "3023-08-30"
	} else {
		expiryDate = user.Expiry
	}
	userQuery := fmt.Sprintf("INSERT INTO %s (user_Id, segment_name, added, expiry) VALUES ($1, $2, $3, $4)", usersTable)
	_, err := r.db.Exec(userQuery, user.UserId, user.SegmentName, additionTime, expiryDate)

	if err != nil {
		return err
	}
	return nil
}

func (r *CreateUserPostgres) DeleteUser(userId int, segmentName string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND segment_name = $2", usersTable)
	_, err := r.db.Exec(query, userId, segmentName)
	return err
}
