package repository

import (
	avito "AvitoTask"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (r *SegmentPostgres) CreateSegment(segment avito.Segment) (bool, error) {
	fmt.Println("CreateSegment_Postgres!!!")
	tx, err := r.db.Begin()
	if err != nil {
		return false, err
	}
	CreateSegmentQuery := fmt.Sprintf("INSERT INTO %s (segment_name) VALUES ($1)", segmentsTable)
	_, err = tx.Exec(CreateSegmentQuery, segment.SegmentName)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	err = tx.Commit()

	if err != nil {
		return false, err
	}

	return true, nil
}
