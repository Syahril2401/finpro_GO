package repository

import (
	"context"
	"finpro/model"

	"gorm.io/gorm"
)

type TargetRepository interface {
	Create(ctx context.Context, t *model.Target) error
	GetByID(ctx context.Context, id, userID string) (*model.Target, error)
	GetByWeek(ctx context.Context, userID string, weekStart, weekEnd string) ([]model.Target, error)
	GetAllByUser(ctx context.Context, userID string) ([]model.Target, error)
	Update(ctx context.Context, t *model.Target) error
	Delete(ctx context.Context, id, userID string) error
	CountByUser(ctx context.Context, userID string) (int64, int64, error)
	CountSubtasksByUser(ctx context.Context, userID string) (int64, int64, error)
	GetPrimaryFocus(ctx context.Context, userID string, weekStart, weekEnd string) (string, error)
	SearchByUser(ctx context.Context, userID, query string) ([]model.Target, error)
	GetCompletedDatesByRange(ctx context.Context, userID string, start, end string) ([]string, error)

	// Subtask operations
	CreateSubtask(ctx context.Context, s *model.Subtask) error
	GetSubtaskByID(ctx context.Context, id, userID string) (*model.Subtask, error)
	UpdateSubtask(ctx context.Context, s *model.Subtask) error
	DeleteSubtask(ctx context.Context, id, userID string) error
	GetSubtasksByTarget(ctx context.Context, targetID string) ([]model.Subtask, error)
	CountSubtasksByTarget(ctx context.Context, targetID string) (int64, int64, error)
}

type targetRepository struct{ db *gorm.DB }

func NewTargetRepository(db *gorm.DB) TargetRepository {
	return &targetRepository{db}
}

func (r *targetRepository) Create(ctx context.Context, t *model.Target) error {
	return r.db.WithContext(ctx).Create(t).Error
}

func (r *targetRepository) GetByID(ctx context.Context, id, userID string) (*model.Target, error) {
	var t model.Target
	err := r.db.WithContext(ctx).Preload("Subtasks").Where("target_id = ? AND user_id = ?", id, userID).First(&t).Error
	return &t, err
}

func (r *targetRepository) GetByWeek(ctx context.Context, userID string, weekStart, weekEnd string) ([]model.Target, error) {
	var targets []model.Target
	err := r.db.WithContext(ctx).Preload("Subtasks").
		Where("user_id = ? AND ((due_date >= ? AND due_date <= ?) OR due_date IS NULL)", userID, weekStart, weekEnd).
		Order("created_at DESC").Find(&targets).Error
	return targets, err
}

func (r *targetRepository) GetAllByUser(ctx context.Context, userID string) ([]model.Target, error) {
	var targets []model.Target
	err := r.db.WithContext(ctx).Preload("Subtasks").Where("user_id = ?", userID).Order("created_at DESC").Find(&targets).Error
	return targets, err
}

func (r *targetRepository) Update(ctx context.Context, t *model.Target) error {
	return r.db.WithContext(ctx).Save(t).Error
}

func (r *targetRepository) Delete(ctx context.Context, id, userID string) error {
	return r.db.WithContext(ctx).Where("target_id = ? AND user_id = ?", id, userID).Delete(&model.Target{}).Error
}

func (r *targetRepository) CountByUser(ctx context.Context, userID string) (int64, int64, error) {
	var total, completed int64
	err := r.db.WithContext(ctx).Model(&model.Target{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return 0, 0, err
	}
	err = r.db.WithContext(ctx).Model(&model.Target{}).Where("user_id = ? AND status = 'completed'", userID).Count(&completed).Error
	return total, completed, err
}

func (r *targetRepository) CountSubtasksByUser(ctx context.Context, userID string) (int64, int64, error) {
	var total, completed int64
	err := r.db.WithContext(ctx).Model(&model.Subtask{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return 0, 0, err
	}
	err = r.db.WithContext(ctx).Model(&model.Subtask{}).Where("user_id = ? AND is_completed = 1", userID).Count(&completed).Error
	return total, completed, err
}

func (r *targetRepository) GetPrimaryFocus(ctx context.Context, userID string, weekStart, weekEnd string) (string, error) {
	var result struct{ FocusDimension string }
	err := r.db.WithContext(ctx).Model(&model.Target{}).
		Select("focus_dimension").
		Where("user_id = ? AND ((due_date >= ? AND due_date <= ?) OR due_date IS NULL)", userID, weekStart, weekEnd).
		Group("focus_dimension").
		Order("COUNT(*) DESC").
		Limit(1).
		Scan(&result).Error
	if result.FocusDimension == "" {
		return "No focus yet", err
	}
	return result.FocusDimension, err
}

func (r *targetRepository) SearchByUser(ctx context.Context, userID, query string) ([]model.Target, error) {
	var targets []model.Target
	like := "%" + query + "%"
	err := r.db.WithContext(ctx).Preload("Subtasks").
		Where("user_id = ? AND (title LIKE ? OR description LIKE ?)", userID, like, like).
		Order("created_at DESC").Limit(10).Find(&targets).Error
	return targets, err
}

func (r *targetRepository) GetCompletedDatesByRange(ctx context.Context, userID string, start, end string) ([]string, error) {
	var dates []string
	err := r.db.WithContext(ctx).Model(&model.Subtask{}).
		Select("DISTINCT DATE(completed_at)").
		Where("user_id = ? AND is_completed = 1 AND completed_at >= ? AND completed_at <= ?", userID, start, end).
		Scan(&dates).Error
	return dates, err
}

// --- Subtask operations ---

func (r *targetRepository) CreateSubtask(ctx context.Context, s *model.Subtask) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *targetRepository) GetSubtaskByID(ctx context.Context, id, userID string) (*model.Subtask, error) {
	var s model.Subtask
	err := r.db.WithContext(ctx).Where("subtask_id = ? AND user_id = ?", id, userID).First(&s).Error
	return &s, err
}

func (r *targetRepository) UpdateSubtask(ctx context.Context, s *model.Subtask) error {
	return r.db.WithContext(ctx).Save(s).Error
}

func (r *targetRepository) DeleteSubtask(ctx context.Context, id, userID string) error {
	return r.db.WithContext(ctx).Where("subtask_id = ? AND user_id = ?", id, userID).Delete(&model.Subtask{}).Error
}

func (r *targetRepository) GetSubtasksByTarget(ctx context.Context, targetID string) ([]model.Subtask, error) {
	var subtasks []model.Subtask
	err := r.db.WithContext(ctx).Where("target_id = ?", targetID).Order("created_at ASC").Find(&subtasks).Error
	return subtasks, err
}

func (r *targetRepository) CountSubtasksByTarget(ctx context.Context, targetID string) (int64, int64, error) {
	var total, completed int64
	err := r.db.WithContext(ctx).Model(&model.Subtask{}).Where("target_id = ?", targetID).Count(&total).Error
	if err != nil {
		return 0, 0, err
	}
	err = r.db.WithContext(ctx).Model(&model.Subtask{}).Where("target_id = ? AND is_completed = 1", targetID).Count(&completed).Error
	return total, completed, err
}
