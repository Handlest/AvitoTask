package service

import (
	avito "AvitoTask"
	"AvitoTask/pkg/repository"
	"fmt"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(segment avito.Segment) (bool, error) {
	fmt.Println("CreateSegment from service")
	return s.repo.CreateSegment(segment)
}
