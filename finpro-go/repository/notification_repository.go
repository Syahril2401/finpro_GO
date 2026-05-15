package repository

import (
	"context"
	"finpro/model"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(ctx context.Context, n *model.Notification) error
	GetByUser(ctx context.Context, userID string) ([]model.Notification, error)
	MarkRead(ctx context.Context, id, userID string) error
	MarkAllRead(ctx context.Context, userID string) error
	CountUnread(ctx context.Context, userID string) (int64, error)
}

type notificationRepository struct{ db *gorm.DB }

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db}
}

func (r *notificationRepository) Create(ctx context.Context, n *model.Notification) error {
	return r.db.WithContext(ctx).Create(n).Error
}

func (r *notificationRepository) GetByUser(ctx context.Context, userID string) ([]model.Notification, error) {
	var notifs []model.Notification
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Limit(50).Find(&notifs).Error
	return notifs, err
}

func (r *notificationRepository) MarkRead(ctx context.Context, id, userID string) error {
	return r.db.WithContext(ctx).Model(&model.Notification{}).
		Where("notification_id = ? AND user_id = ?", id, userID).
		Update("is_read", true).Error
}

func (r *notificationRepository) MarkAllRead(ctx context.Context, userID string) error {
	return r.db.WithContext(ctx).Model(&model.Notification{}).
		Where("user_id = ? AND is_read = 0", userID).
		Update("is_read", true).Error
}

func (r *notificationRepository) CountUnread(ctx context.Context, userID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Notification{}).Where("user_id = ? AND is_read = 0", userID).Count(&count).Error
	return count, err
}
