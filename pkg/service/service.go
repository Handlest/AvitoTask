package service

import (
	avito "AvitoTask"
	"AvitoTask/pkg/repository"
)

type User interface {
	CreateUser(user avito.User) (int, error)
}

type Segment interface {
}
type Operation interface {
}
type Service struct {
	User
	Segment
	Operation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{User: NewCreateService(repos.User)}
}
