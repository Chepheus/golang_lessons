package service

import (
	"golang_lessons/custom_errors"
	"golang_lessons/domain"
	"golang_lessons/dto"
	"golang_lessons/logger"
)

type UserService interface {
	Login(dto.LoginRequest) (*string, *custom_errors.AppErrors)
}

type DefaultUserService struct {
	repoU domain.UserRepository
	repoL domain.LoginRepository
}

func (s DefaultUserService) Login(request dto.LoginRequest) (*string, *custom_errors.AppErrors) {
	logger.Info("test")
	login, err := s.repoL.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}

	return token, nil
}

func NewUserService(repoU domain.UserRepository, repoL domain.LoginRepository) UserService {
	return DefaultUserService{
		repoU: repoU,
		repoL: repoL,
	}
}
