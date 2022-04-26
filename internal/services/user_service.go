package services

import (
	"crud-task/internal/repositories"
	"crud-task/internal/requests"
	"errors"
)

type UserService struct {
	repository *repositories.UserRepository
}

func (service *UserService) Create(request *requests.UserRequest) error {
	err := service.repository.Create(request.User)

	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (service *UserService) Update(request *requests.UserRequest) error {
	err := service.repository.Update(request.User)

	if err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func (service *UserService) Delete(request *requests.UserRequest) error {
	err := service.repository.Delete(request.User)

	if err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

func NewUserService() *UserService {
	return &UserService{
		repository: repositories.NewUserRepository(),
	}
}
