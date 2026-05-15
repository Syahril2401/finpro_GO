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
	GetTargetStats(userID string) (int64, int64, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db}
}

func (r *dashboardRepository) GetTodaySchedules(userID string) ([]model.Schedule, error) {
	var schedules []model.Schedule
	today := time.Now().Format("2006-01-02")
	err := r.db.Where("user_id = ? AND date = ?", userID, today).Order("start_time ASC").Find(&schedules).Error
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
	var total *int
	err := r.db.Model(&model.Schedule{}).
		Select("COALESCE(SUM(duration_minutes), 0)").
		Where("user_id = ? AND status = 'completed'", userID).
		Scan(&total).Error
	if err != nil || total == nil {
		return 0, err
	}
	return float64(*total) / 60.0, nil
}

func (r *dashboardRepository) GetTargetStats(userID string) (int64, int64, error) {
	var total int64
	var completed int64
	err := r.db.Model(&model.Target{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return 0, 0, err
	}
	err = r.db.Model(&model.Target{}).Where("user_id = ? AND status = ?", userID, "completed").Count(&completed).Error
	return total, completed, err
}
