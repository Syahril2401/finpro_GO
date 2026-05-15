package service

import (
	"context"
	"finpro/model"
	"finpro/repository"
	"time"

	"github.com/google/uuid"
)

type TargetService interface {
	Create(ctx context.Context, userID string, req CreateTargetReq) (*model.Target, error)
	GetByID(ctx context.Context, userID, id string) (*model.Target, error)
	GetByWeek(ctx context.Context, userID, weekParam string) ([]model.Target, error)
	GetAll(ctx context.Context, userID string) ([]model.Target, error)
	Update(ctx context.Context, userID, id string, req UpdateTargetReq) (*model.Target, error)
	Delete(ctx context.Context, userID, id string) error
	Search(ctx context.Context, userID, query string) ([]model.Target, error)

	CreateSubtask(ctx context.Context, userID, targetID string, req CreateSubtaskReq) (*model.Subtask, error)
	UpdateSubtask(ctx context.Context, userID, subtaskID string, req UpdateSubtaskReq) (*model.Subtask, error)
	DeleteSubtask(ctx context.Context, userID, subtaskID string) error
	ToggleSubtask(ctx context.Context, userID, subtaskID string) (*model.Subtask, error)

	GetSummary(ctx context.Context, userID string, weekStart, weekEnd string) (map[string]interface{}, error)
}

type targetService struct {
	repo repository.TargetRepository
}

func NewTargetService(repo repository.TargetRepository) TargetService {
	return &targetService{repo}
}

type CreateTargetReq struct {
	Title          string           `json:"title" binding:"required"`
	Description    string           `json:"description"`
	FocusDimension string           `json:"focus_dimension"`
	Priority       string           `json:"priority"`
	DueDate        *string          `json:"due_date"`
	Subtasks       []SubtaskInput   `json:"subtasks"`
}

type SubtaskInput struct {
	Title string `json:"title"`
}

type UpdateTargetReq struct {
	Title          *string `json:"title"`
	Description    *string `json:"description"`
	FocusDimension *string `json:"focus_dimension"`
	Priority       *string `json:"priority"`
	DueDate        *string `json:"due_date"`
	Status         *string `json:"status"`
}

type CreateSubtaskReq struct {
	Title string `json:"title" binding:"required"`
}

type UpdateSubtaskReq struct {
	Title *string `json:"title"`
}

func (s *targetService) Create(ctx context.Context, userID string, req CreateTargetReq) (*model.Target, error) {
	if req.FocusDimension == "" {
		req.FocusDimension = "General"
	}
	if req.Priority == "" {
		req.Priority = "medium"
	}
	now := time.Now()
	_, week := now.ISOWeek()
	year := now.Year()
	target := &model.Target{
		TargetID:       uuid.New().String(),
		UserID:         userID,
		Title:          req.Title,
		Description:    req.Description,
		FocusDimension: req.FocusDimension,
		Priority:       req.Priority,
		DueDate:        req.DueDate,
		Status:         "not_started",
		Progress:       0,
		WeekNumber:     &week,
		Year:           &year,
	}
	if err := s.repo.Create(ctx, target); err != nil {
		return nil, err
	}
	// Create subtasks if provided
	for _, st := range req.Subtasks {
		if st.Title == "" {
			continue
		}
		subtask := &model.Subtask{
			SubtaskID: uuid.New().String(),
			TargetID:  target.TargetID,
			UserID:    userID,
			Title:     st.Title,
		}
		s.repo.CreateSubtask(ctx, subtask)
	}
	return s.repo.GetByID(ctx, target.TargetID, userID)
}

func (s *targetService) GetByID(ctx context.Context, userID, id string) (*model.Target, error) {
	return s.repo.GetByID(ctx, id, userID)
}

func (s *targetService) GetByWeek(ctx context.Context, userID, weekParam string) ([]model.Target, error) {
	var start, end string
	if weekParam != "" {
		t, err := time.Parse("2006-01-02", weekParam)
		if err == nil {
			start, end = repository.WeekBounds(t)
		}
	}
	if start == "" {
		start, end = repository.WeekBounds(time.Now())
	}
	return s.repo.GetByWeek(ctx, userID, start, end)
}

func (s *targetService) GetAll(ctx context.Context, userID string) ([]model.Target, error) {
	return s.repo.GetAllByUser(ctx, userID)
}

func (s *targetService) Update(ctx context.Context, userID, id string, req UpdateTargetReq) (*model.Target, error) {
	t, err := s.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if req.Title != nil {
		t.Title = *req.Title
	}
	if req.Description != nil {
		t.Description = *req.Description
	}
	if req.FocusDimension != nil {
		t.FocusDimension = *req.FocusDimension
	}
	if req.Priority != nil {
		t.Priority = *req.Priority
	}
	if req.DueDate != nil {
		t.DueDate = req.DueDate
	}
	if req.Status != nil {
		t.Status = *req.Status
	}
	err = s.repo.Update(ctx, t)
	return t, err
}

func (s *targetService) Delete(ctx context.Context, userID, id string) error {
	return s.repo.Delete(ctx, id, userID)
}

func (s *targetService) Search(ctx context.Context, userID, query string) ([]model.Target, error) {
	return s.repo.SearchByUser(ctx, userID, query)
}

// --- Subtask operations ---

func (s *targetService) CreateSubtask(ctx context.Context, userID, targetID string, req CreateSubtaskReq) (*model.Subtask, error) {
	subtask := &model.Subtask{
		SubtaskID: uuid.New().String(),
		TargetID:  targetID,
		UserID:    userID,
		Title:     req.Title,
	}
	if err := s.repo.CreateSubtask(ctx, subtask); err != nil {
		return nil, err
	}
	s.recalcProgress(ctx, targetID, userID)
	return subtask, nil
}

func (s *targetService) UpdateSubtask(ctx context.Context, userID, subtaskID string, req UpdateSubtaskReq) (*model.Subtask, error) {
	sub, err := s.repo.GetSubtaskByID(ctx, subtaskID, userID)
	if err != nil {
		return nil, err
	}
	if req.Title != nil {
		sub.Title = *req.Title
	}
	err = s.repo.UpdateSubtask(ctx, sub)
	return sub, err
}

func (s *targetService) DeleteSubtask(ctx context.Context, userID, subtaskID string) error {
	sub, err := s.repo.GetSubtaskByID(ctx, subtaskID, userID)
	if err != nil {
		return err
	}
	targetID := sub.TargetID
	if err := s.repo.DeleteSubtask(ctx, subtaskID, userID); err != nil {
		return err
	}
	s.recalcProgress(ctx, targetID, userID)
	return nil
}

func (s *targetService) ToggleSubtask(ctx context.Context, userID, subtaskID string) (*model.Subtask, error) {
	sub, err := s.repo.GetSubtaskByID(ctx, subtaskID, userID)
	if err != nil {
		return nil, err
	}
	sub.IsCompleted = !sub.IsCompleted
	if sub.IsCompleted {
		now := time.Now()
		sub.CompletedAt = &now
	} else {
		sub.CompletedAt = nil
	}
	err = s.repo.UpdateSubtask(ctx, sub)
	if err != nil {
		return nil, err
	}
	s.recalcProgress(ctx, sub.TargetID, userID)
	return sub, nil
}

func (s *targetService) recalcProgress(ctx context.Context, targetID, userID string) {
	total, completed, err := s.repo.CountSubtasksByTarget(ctx, targetID)
	if err != nil {
		return
	}
	target, err := s.repo.GetByID(ctx, targetID, userID)
	if err != nil {
		return
	}
	if total == 0 {
		target.Progress = 0
		target.Status = "not_started"
	} else {
		target.Progress = int(completed * 100 / total)
		if target.Progress == 0 {
			target.Status = "not_started"
		} else if target.Progress >= 100 {
			target.Progress = 100
			target.Status = "completed"
		} else {
			target.Status = "in_progress"
		}
	}
	s.repo.Update(ctx, target)
}

func (s *targetService) GetSummary(ctx context.Context, userID string, weekStart, weekEnd string) (map[string]interface{}, error) {
	totalTargets, completedTargets, _ := s.repo.CountByUser(ctx, userID)
	totalSubtasks, completedSubtasks, _ := s.repo.CountSubtasksByUser(ctx, userID)
	primaryFocus, _ := s.repo.GetPrimaryFocus(ctx, userID, weekStart, weekEnd)

	var completionRate float64
	if totalSubtasks > 0 {
		completionRate = float64(completedSubtasks) / float64(totalSubtasks) * 100
	} else if totalTargets > 0 {
		completionRate = float64(completedTargets) / float64(totalTargets) * 100
	}

	return map[string]interface{}{
		"total_targets":      totalTargets,
		"completed_targets":  completedTargets,
		"total_subtasks":     totalSubtasks,
		"completed_subtasks": completedSubtasks,
		"completion_rate":    completionRate,
		"primary_focus":      primaryFocus,
	}, nil
}
