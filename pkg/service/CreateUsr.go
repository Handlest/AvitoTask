package service

import (
	avito "AvitoTask"
	"AvitoTask/pkg/repository"
	"fmt"
)

type CreateService struct {
	repo repository.User
}

func NewCreateService(repo repository.User) *CreateService {
	return &CreateService{repo: repo}
}

func (s *CreateService) CreateUser(user avito.User) (int, error) {
	fmt.Println("CreateUSR from service")
	return s.repo.CreateUser(user)
}
