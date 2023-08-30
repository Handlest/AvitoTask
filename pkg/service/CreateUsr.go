package service

import (
	avito "AvitoTask"
	"AvitoTask/pkg/repository"
)

type CreateService struct {
	repo repository.User
}

func NewCreateService(repo repository.User) *CreateService {
	return &CreateService{repo: repo}
}

func (s *CreateService) CreateUser(user avito.User) error {
	return s.repo.CreateUser(user)
}

func (s *CreateService) DeleteUser(userId int, segmentName string) error {
	return s.repo.DeleteUser(userId, segmentName)
}
