package service

import (
	"errors"
	"finpro/model"
	"finpro/repository"
	"finpro/utils"
)

type AuthService interface {
	Register(name, email, password string) (string, error)
	Login(email, password string) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Register(name, email, password string) (string, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     "user",
	}

	if err := s.userRepo.Create(user); err != nil {
		return "", err
	}

	return utils.GenerateToken(user.ID, user.Email, user.Role)
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	return utils.GenerateToken(user.ID, user.Email, user.Role)
}
