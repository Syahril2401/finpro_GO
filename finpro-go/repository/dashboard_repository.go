package repository

import (
	"finpro/model"
	"time"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	GetTodaySchedules(userID string) ([]model.Schedule, error)
	GetActiveTargets(userID string) ([]model.Target, error)
	GetLastAILog(userID string) (*model.AILog, error)
	GetRecentAILogs(userID string, limit int) ([]model.AILog, error)
	GetTotalFocusSessions(userID string) (int64, error)
	GetTotalDeepWorkHours(userID string) (float64, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db}
}

func (r *dashboardRepository) GetTodaySchedules(userID string) ([]model.Schedule, error) {
	var schedules []model.Schedule
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := r.db.Where("user_id = ? AND start_time >= ? AND start_time < ?", userID, startOfDay, endOfDay).Find(&schedules).Error
	return schedules, err
}

func (r *dashboardRepository) GetActiveTargets(userID string) ([]model.Target, error) {
	var targets []model.Target
	err := r.db.Where("user_id = ? AND status != ?", userID, "completed").Find(&targets).Error
	return targets, err
}

func (r *dashboardRepository) GetLastAILog(userID string) (*model.AILog, error) {
	var log model.AILog
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").First(&log).Error
	return &log, err
}

func (r *dashboardRepository) GetRecentAILogs(userID string, limit int) ([]model.AILog, error) {
	var logs []model.AILog
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Limit(limit).Find(&logs).Error
	return logs, err
}

func (r *dashboardRepository) GetTotalFocusSessions(userID string) (int64, error) {
	var count int64
	err := r.db.Model(&model.Schedule{}).Where("user_id = ? AND status = ?", userID, "done").Count(&count).Error
	return count, err
}

func (r *dashboardRepository) GetTotalDeepWorkHours(userID string) (float64, error) {
	var schedules []model.Schedule
	err := r.db.Where("user_id = ? AND status = ? AND end_time IS NOT NULL", userID, "done").Find(&schedules).Error
	if err != nil {
		return 0, err
	}

	var total float64
	for _, s := range schedules {
		if s.EndTime != nil {
			total += s.EndTime.Sub(s.StartTime).Hours()
		}
	}
	return total, nil
}
