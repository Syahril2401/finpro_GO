package service

import (
	"context"
	"errors"
	"finpro/model"
	"finpro/repository"
	"time"

	"github.com/google/uuid"
)

type NoteService interface {
	CreateNote(ctx context.Context, userID string, req CreateNoteRequest) (*model.Note, error)
	GetNote(ctx context.Context, userID string, noteID string) (*model.Note, error)
	GetAllNotes(ctx context.Context, userID string) ([]model.Note, error)
	SearchNotes(ctx context.Context, userID string, query string) ([]model.Note, error)
	UpdateNote(ctx context.Context, userID string, noteID string, req UpdateNoteRequest) (*model.Note, error)
	DeleteNote(ctx context.Context, userID string, noteID string) error
	TogglePin(ctx context.Context, userID string, noteID string) (*model.Note, error)
	ToggleArchive(ctx context.Context, userID string, noteID string) (*model.Note, error)
}

type noteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteService{repo: repo}
}

type CreateNoteRequest struct {
	Title            string `json:"title"`
	ContentJSON      string `json:"content_json"`
	ContentText      string `json:"content_text"`
	FocusDimension   string `json:"focus_dimension"`
	Tags             string `json:"tags"`
	Mood             string `json:"mood"`
	ConfidenceLevel  *int   `json:"confidence_level"`
	PlannerSessionID *string `json:"planner_session_id"`
	TargetID         *string `json:"target_id"`
}

type UpdateNoteRequest struct {
	Title            *string `json:"title"`
	ContentJSON      *string `json:"content_json"`
	ContentText      *string `json:"content_text"`
	FocusDimension   *string `json:"focus_dimension"`
	Tags             *string `json:"tags"`
	Mood             *string `json:"mood"`
	ConfidenceLevel  *int    `json:"confidence_level"`
	PlannerSessionID *string `json:"planner_session_id"`
	TargetID         *string `json:"target_id"`
}

func (s *noteService) CreateNote(ctx context.Context, userID string, req CreateNoteRequest) (*model.Note, error) {
	if req.Title == "" {
		req.Title = "Untitled"
	}
	if req.FocusDimension == "" {
		req.FocusDimension = "General"
	}

	note := &model.Note{
		NoteID:           uuid.New().String(),
		UserID:           userID,
		Title:            req.Title,
		ContentJSON:      req.ContentJSON,
		ContentText:      req.ContentText,
		FocusDimension:   req.FocusDimension,
		Tags:             req.Tags,
		Mood:             req.Mood,
		ConfidenceLevel:  req.ConfidenceLevel,
		PlannerSessionID: req.PlannerSessionID,
		TargetID:         req.TargetID,
		IsPinned:         false,
		IsArchived:       false,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := s.repo.Create(ctx, note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s *noteService) GetNote(ctx context.Context, userID string, noteID string) (*model.Note, error) {
	return s.repo.GetByID(ctx, noteID, userID)
}

func (s *noteService) GetAllNotes(ctx context.Context, userID string) ([]model.Note, error) {
	return s.repo.GetAllByUserID(ctx, userID)
}

func (s *noteService) SearchNotes(ctx context.Context, userID string, query string) ([]model.Note, error) {
	if query == "" {
		return s.repo.GetAllByUserID(ctx, userID)
	}
	return s.repo.SearchByUserID(ctx, userID, query)
}

func (s *noteService) UpdateNote(ctx context.Context, userID string, noteID string, req UpdateNoteRequest) (*model.Note, error) {
	note, err := s.repo.GetByID(ctx, noteID, userID)
	if err != nil {
		return nil, err
	}
	if note == nil {
		return nil, errors.New("note not found")
	}

	if req.Title != nil {
		note.Title = *req.Title
	}
	if req.ContentJSON != nil {
		note.ContentJSON = *req.ContentJSON
	}
	if req.ContentText != nil {
		note.ContentText = *req.ContentText
	}
	if req.FocusDimension != nil {
		note.FocusDimension = *req.FocusDimension
	}
	if req.Tags != nil {
		note.Tags = *req.Tags
	}
	if req.Mood != nil {
		note.Mood = *req.Mood
	}
	if req.ConfidenceLevel != nil {
		note.ConfidenceLevel = req.ConfidenceLevel
	}
	if req.PlannerSessionID != nil {
		note.PlannerSessionID = req.PlannerSessionID
	}
	if req.TargetID != nil {
		note.TargetID = req.TargetID
	}

	note.UpdatedAt = time.Now()
	err = s.repo.Update(ctx, note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s *noteService) DeleteNote(ctx context.Context, userID string, noteID string) error {
	return s.repo.Delete(ctx, noteID, userID)
}

func (s *noteService) TogglePin(ctx context.Context, userID string, noteID string) (*model.Note, error) {
	note, err := s.repo.GetByID(ctx, noteID, userID)
	if err != nil {
		return nil, err
	}
	note.IsPinned = !note.IsPinned
	note.UpdatedAt = time.Now()
	err = s.repo.Update(ctx, note)
	return note, err
}

func (s *noteService) ToggleArchive(ctx context.Context, userID string, noteID string) (*model.Note, error) {
	note, err := s.repo.GetByID(ctx, noteID, userID)
	if err != nil {
		return nil, err
	}
	note.IsArchived = !note.IsArchived
	note.UpdatedAt = time.Now()
	err = s.repo.Update(ctx, note)
	return note, err
}
