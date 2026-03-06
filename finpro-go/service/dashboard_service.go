package service

import (
	"finpro/repository"
)

type DashboardService interface {
	GetDashboardData(userID string) (map[string]interface{}, error)
}

type dashboardService struct {
	dashboardRepo  repository.DashboardRepository
	assessmentRepo repository.AssessmentRepository
}

func NewDashboardService(dashRepo repository.DashboardRepository, assessRepo repository.AssessmentRepository) DashboardService {
	return &dashboardService{dashRepo, assessRepo}
}

func (s *dashboardService) GetDashboardData(userID string) (map[string]interface{}, error) {
	schedules, _ := s.dashboardRepo.GetTodaySchedules(userID)
	targets, _ := s.dashboardRepo.GetActiveTargets(userID)
	aiLog, _ := s.dashboardRepo.GetLastAILog(userID)
	lastAssessment, _ := s.assessmentRepo.GetLastSummary(userID)

	data := map[string]interface{}{
		"schedules":       schedules,
		"targets":         targets,
		"last_insight":    aiLog,
		"last_assessment": lastAssessment,
	}

	return data, nil
}
