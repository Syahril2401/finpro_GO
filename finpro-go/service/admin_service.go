package service

import (
	"errors"
	"finpro/model"
	"finpro/repository"
	"finpro/utils"
)

type AdminService interface {
	GetAllUsers() ([]model.User, error)
	GetAllEvaluations() ([]model.Evaluation, error)
	Login(email string, password string) (string, error)
}

func (s *adminService) Login(email string, password string) (string, error) {
	admin, err := s.adminRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, admin.PasswordHash) {
		return "", errors.New("invalid email or password")
	}

	return utils.GenerateToken(admin.AdminID, admin.Email, "admin")
}

type adminService struct {
	adminRepo      repository.AdminRepository
	userRepo       repository.UserRepository
	assessmentRepo repository.AssessmentRepository
}

func NewAdminService(adminRepo repository.AdminRepository, userRepo repository.UserRepository, assessmentRepo repository.AssessmentRepository) AdminService {
	return &adminService{adminRepo, userRepo, assessmentRepo}
}

func (s *adminService) GetAllUsers() ([]model.User, error) {
	return s.userRepo.FindAll()
}

func (s *adminService) GetAllEvaluations() ([]model.Evaluation, error) {
	return s.assessmentRepo.GetAllEvaluations()
}
