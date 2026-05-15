package repository

import (
	"context"
	"finpro/model"

	"gorm.io/gorm"
)

type NoteRepository interface {
	Create(ctx context.Context, note *model.Note) error
	GetByID(ctx context.Context, noteID string, userID string) (*model.Note, error)
	GetAllByUserID(ctx context.Context, userID string) ([]model.Note, error)
	SearchByUserID(ctx context.Context, userID string, query string) ([]model.Note, error)
	Update(ctx context.Context, note *model.Note) error
	Delete(ctx context.Context, noteID string, userID string) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{db: db}
}

func (r *noteRepository) Create(ctx context.Context, note *model.Note) error {
	return r.db.WithContext(ctx).Create(note).Error
}

func (r *noteRepository) GetByID(ctx context.Context, noteID string, userID string) (*model.Note, error) {
	var note model.Note
	err := r.db.WithContext(ctx).Where("note_id = ? AND user_id = ?", noteID, userID).First(&note).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *noteRepository) GetAllByUserID(ctx context.Context, userID string) ([]model.Note, error) {
	var notes []model.Note
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("updated_at desc").Find(&notes).Error
	return notes, err
}

func (r *noteRepository) SearchByUserID(ctx context.Context, userID string, query string) ([]model.Note, error) {
	var notes []model.Note
	likeQuery := "%" + query + "%"
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND (title LIKE ? OR content_text LIKE ? OR tags LIKE ?)", userID, likeQuery, likeQuery, likeQuery).
		Order("updated_at desc").
		Find(&notes).Error
	return notes, err
}

func (r *noteRepository) Update(ctx context.Context, note *model.Note) error {
	return r.db.WithContext(ctx).Save(note).Error
}

func (r *noteRepository) Delete(ctx context.Context, noteID string, userID string) error {
	return r.db.WithContext(ctx).Where("note_id = ? AND user_id = ?", noteID, userID).Delete(&model.Note{}).Error
}
