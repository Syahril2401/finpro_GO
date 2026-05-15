package service

import (
	"encoding/json"
	"fmt"
	"finpro/repository"
)

type DashboardService interface {
	GetDashboardData(userID string) (map[string]interface{}, error)
}

type dashboardService struct {
	dashboardRepo  repository.DashboardRepository
	assessmentRepo repository.AssessmentRepository
	userRepo       repository.UserRepository
}

func NewDashboardService(dashRepo repository.DashboardRepository, assessRepo repository.AssessmentRepository, userRepo repository.UserRepository) DashboardService {
	return &dashboardService{dashRepo, assessRepo, userRepo}
}

func (s *dashboardService) GetDashboardData(userID string) (map[string]interface{}, error) {
	// 1. Fetch User Info
	user, _ := s.userRepo.FindByID(userID)
	userName := "User"
	if user != nil {
		userName = user.Name
	}

	// 2. Fetch Metrics
	focusSessions, _ := s.dashboardRepo.GetTotalFocusSessions(userID)
	deepWorkHours, _ := s.dashboardRepo.GetTotalDeepWorkHours(userID)
	lastAssessment, _ := s.assessmentRepo.GetLastSummary(userID)

	// Default fallbacks matching Result.vue
	consistency := 64
	retention := 88
	profileTitle := "Focused Achiever"
	
	if lastAssessment != nil {
		// Parse AI JSON for consistency, retention, and profile title
		if lastAssessment.CategoryResult != nil {
			var aiData map[string]interface{}
			if err := json.Unmarshal([]byte(*lastAssessment.CategoryResult), &aiData); err == nil {
				if title, ok := aiData["profile_title"].(string); ok {
					profileTitle = title
				}
				if val, ok := aiData["consistency_score"].(float64); ok {
					consistency = int(val)
				}
				if val, ok := aiData["retention_score"].(float64); ok {
					retention = int(val)
				}
			}
		}
	}

	// 3. Fetch Today's Sessions
	schedules, _ := s.dashboardRepo.GetTodaySchedules(userID)
	formattedSessions := []map[string]interface{}{}
	for _, sch := range schedules {
		duration := fmt.Sprintf("%dm", sch.DurationMinutes)
		formattedSessions = append(formattedSessions, map[string]interface{}{
			"title":    sch.Title,
			"time":     sch.StartTime,
			"duration": duration,
			"status":   sch.Status,
		})
	}

	// 4. Fetch AI Messages (Logs)
	aiLogs, _ := s.dashboardRepo.GetRecentAILogs(userID, 5)
	aiMessages := []map[string]interface{}{}
	for _, log := range aiLogs {
		var aiData map[string]interface{}
		content := "Your consistency in planning is improving. Keep up the good work!"
		if err := json.Unmarshal([]byte(string(log.AIOutput)), &aiData); err == nil {
			if strategy, ok := aiData["ai_strategy"].(map[string]interface{}); ok {
				if desc, ok := strategy["desc"].(string); ok {
					content = desc
				}
			}
		}
		aiMessages = append(aiMessages, map[string]interface{}{
			"timestamp": log.CreatedAt.Format("15:04"),
			"content":   content,
		})
	}

	totalTargets, completedTargets, _ := s.dashboardRepo.GetTargetStats(userID)
	taskEfficiency := 87 // fallback
	if totalTargets > 0 {
		taskEfficiency = int((completedTargets * 100) / totalTargets)
	}

	data := map[string]interface{}{
		"user": map[string]interface{}{
			"name":   userName,
			"status": "online",
		},
		"metrics": map[string]interface{}{
			"focus_sessions":  focusSessions,
			"task_efficiency": taskEfficiency, 
			"deep_work_hours": fmt.Sprintf("%.1f", deepWorkHours),
			"consistency":     consistency,
			"retention":       retention,
		},
		"focus_sessions":  formattedSessions,
		"ai_messages":     aiMessages,
		"profile_title":   profileTitle,
	}

	return data, nil
}
