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
