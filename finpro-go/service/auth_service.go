package service

import (
	"errors"
	"finpro/model"
	"finpro/repository"
	"finpro/utils"

	"github.com/google/uuid"
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
		UserID:       uuid.New().String(),
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
		IsActive:     true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return "", err
	}

	// Default role logic (since Role field is gone, we can set it to "user" in token for now)
	return utils.GenerateToken(user.UserID, user.Email, "user")
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid email or password")
	}

	return utils.GenerateToken(user.UserID, user.Email, "user")
}
