package service

import (
	avito "AvitoTask"
	"AvitoTask/pkg/repository"
)

type User interface {
	CreateUser(user avito.User) error
	DeleteUser(userId int, segmentName string) error
	GetOperations(user avito.UserInfo) ([]avito.Operation, error)
}

type Segment interface {
	CreateSegment(segment avito.Segment) (bool, error)
	GetSegments(userId int) ([]avito.Segment, error)
	DeleteSegment(segmentName string) error
}
type Operation interface {
}
type Service struct {
	User
	Segment
	Operation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:    NewCreateService(repos.User),
		Segment: NewSegmentService(repos.Segment),
	}
}
