package service

import (
	"finpro/model"
	"finpro/repository"
)

type AdminService interface {
	GetAllUsers() ([]model.User, error)
	GetAllEvaluations() ([]model.Evaluation, error)
}

type adminService struct {
	userRepo       repository.UserRepository
	assessmentRepo repository.AssessmentRepository
}

func NewAdminService(userRepo repository.UserRepository, assessmentRepo repository.AssessmentRepository) AdminService {
	return &adminService{userRepo, assessmentRepo}
}

func (s *adminService) GetAllUsers() ([]model.User, error) {
	return s.userRepo.FindAll()
}

func (s *adminService) GetAllEvaluations() ([]model.Evaluation, error) {
	return s.assessmentRepo.GetAllEvaluations()
}
