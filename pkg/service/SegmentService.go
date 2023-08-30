package service

import (
	avito "AvitoTask"
	"AvitoTask/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(segment avito.Segment) (bool, error) {
	return s.repo.CreateSegment(segment)
}

func (s *SegmentService) GetSegments(userId int) ([]avito.Segment, error) {
	return s.repo.GetSegments(userId)
}

func (s *SegmentService) DeleteSegment(segmentName string) error {
	return s.repo.DeleteSegment(segmentName)
}
