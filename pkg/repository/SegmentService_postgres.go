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
	CreateSegmentQuery := fmt.Sprintf("INSERT INTO %s (segment_name) VALUES ($1)", segmentsTable)
	_, err := r.db.Exec(CreateSegmentQuery, segment.SegmentName)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *SegmentPostgres) GetSegments(userId int) ([]avito.Segment, error) {
	var segments []avito.Segment
	query := fmt.Sprintf("SELECT segment_name FROM %s WHERE user_id = $1", usersTable)
	err := r.db.Select(&segments, query, userId)

	return segments, err
}

func (r *SegmentPostgres) DeleteSegment(segmentName string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE segment_name = $1", segmentsTable)
	_, err := r.db.Exec(query, segmentName)
	return err
}
