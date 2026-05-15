package repository

import (
	"context"
	"finpro/model"
	"time"

	"gorm.io/gorm"
)

type PlannerRepository interface {
	Create(ctx context.Context, s *model.Schedule) error
	GetByID(ctx context.Context, id, userID string) (*model.Schedule, error)
	GetByWeek(ctx context.Context, userID string, startDate, endDate string) ([]model.Schedule, error)
	GetByDate(ctx context.Context, userID string, date string) ([]model.Schedule, error)
	Update(ctx context.Context, s *model.Schedule) error
	Delete(ctx context.Context, id, userID string) error
	CountCompletedThisWeek(ctx context.Context, userID string, weekStart, weekEnd string) (int64, error)
	SumDurationCompletedThisWeek(ctx context.Context, userID string, weekStart, weekEnd string) (int, error)
	GetCompletedDatesByRange(ctx context.Context, userID string, start, end string) ([]string, error)
	SearchByUser(ctx context.Context, userID, query string) ([]model.Schedule, error)
}

type plannerRepository struct{ db *gorm.DB }

func NewPlannerRepository(db *gorm.DB) PlannerRepository {
	return &plannerRepository{db}
}

func (r *plannerRepository) Create(ctx context.Context, s *model.Schedule) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *plannerRepository) GetByID(ctx context.Context, id, userID string) (*model.Schedule, error) {
	var s model.Schedule
	err := r.db.WithContext(ctx).Where("schedule_id = ? AND user_id = ?", id, userID).First(&s).Error
	return &s, err
}

func (r *plannerRepository) GetByWeek(ctx context.Context, userID string, startDate, endDate string) ([]model.Schedule, error) {
	var schedules []model.Schedule
	err := r.db.WithContext(ctx).Where("user_id = ? AND date >= ? AND date <= ?", userID, startDate, endDate).
		Order("date ASC, start_time ASC").Find(&schedules).Error
	return schedules, err
}

func (r *plannerRepository) GetByDate(ctx context.Context, userID string, date string) ([]model.Schedule, error) {
	var schedules []model.Schedule
	err := r.db.WithContext(ctx).Where("user_id = ? AND date = ?", userID, date).
		Order("start_time ASC").Find(&schedules).Error
	return schedules, err
}

func (r *plannerRepository) Update(ctx context.Context, s *model.Schedule) error {
	return r.db.WithContext(ctx).Save(s).Error
}

func (r *plannerRepository) Delete(ctx context.Context, id, userID string) error {
	return r.db.WithContext(ctx).Where("schedule_id = ? AND user_id = ?", id, userID).Delete(&model.Schedule{}).Error
}

func (r *plannerRepository) CountCompletedThisWeek(ctx context.Context, userID string, weekStart, weekEnd string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Schedule{}).
		Where("user_id = ? AND status = 'completed' AND date >= ? AND date <= ?", userID, weekStart, weekEnd).
		Count(&count).Error
	return count, err
}

func (r *plannerRepository) SumDurationCompletedThisWeek(ctx context.Context, userID string, weekStart, weekEnd string) (int, error) {
	var total *int
	err := r.db.WithContext(ctx).Model(&model.Schedule{}).
		Select("COALESCE(SUM(duration_minutes), 0)").
		Where("user_id = ? AND status = 'completed' AND date >= ? AND date <= ?", userID, weekStart, weekEnd).
		Scan(&total).Error
	if total == nil {
		return 0, err
	}
	return *total, err
}

func (r *plannerRepository) GetCompletedDatesByRange(ctx context.Context, userID string, start, end string) ([]string, error) {
	var dates []string
	err := r.db.WithContext(ctx).Model(&model.Schedule{}).
		Select("DISTINCT date").
		Where("user_id = ? AND status = 'completed' AND date >= ? AND date <= ?", userID, start, end).
		Scan(&dates).Error
	return dates, err
}

func (r *plannerRepository) SearchByUser(ctx context.Context, userID, query string) ([]model.Schedule, error) {
	var schedules []model.Schedule
	like := "%" + query + "%"
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND (title LIKE ? OR description LIKE ?)", userID, like, like).
		Order("date DESC").Limit(10).Find(&schedules).Error
	return schedules, err
}

// Helper
func WeekBounds(t time.Time) (string, string) {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	monday := t.AddDate(0, 0, -(weekday - 1))
	sunday := monday.AddDate(0, 0, 6)
	return monday.Format("2006-01-02"), sunday.Format("2006-01-02")
}
