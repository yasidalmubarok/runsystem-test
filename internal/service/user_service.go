package service

import (
	"runsystem-test/internal/dto"
	"runsystem-test/internal/entity"
	"runsystem-test/internal/helper"
	"runsystem-test/internal/repository"
)

type UserService interface {
	CreateUser(payLoad *dto.UserRequest) (*dto.UserResponse, helper.ErrorInterface)
	GetUserByID(id int) (*dto.UserResponse, helper.ErrorInterface)
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(payLoad *dto.UserRequest) (*dto.UserResponse, helper.ErrorInterface) {
	existingUser, err := s.userRepo.GetUserByName(payLoad.Name)
	if err == nil && existingUser != nil {
		return nil, helper.NewConflictError("user with the same name already exists")
	}

	userEntity := &entity.UserEntity{
		Name:    payLoad.Name,
		Hobbies: payLoad.Hobbies,
	}

	savedUser, err := s.userRepo.SaveUser(userEntity)
	if err != nil {
		return nil, err
	}

	userResponse := &dto.UserResponse{
		ID:      savedUser.ID,
		Name:    savedUser.Name,
		Hobbies: savedUser.Hobbies,
	}

	return userResponse, nil
}

func (s *userService) GetUserByID(id int) (*dto.UserResponse, helper.ErrorInterface) {
	userEntity, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	userResponse := &dto.UserResponse{
		ID:      userEntity.ID,
		Name:    userEntity.Name,
		Hobbies: userEntity.Hobbies,
	}

	return userResponse, nil
}