package service

import (
	"context"
	"finpro/model"
	"finpro/repository"
	"time"

	"github.com/google/uuid"
)

type PlannerService interface {
	Create(ctx context.Context, userID string, req CreateScheduleReq) (*model.Schedule, error)
	GetByID(ctx context.Context, userID, id string) (*model.Schedule, error)
	GetByWeek(ctx context.Context, userID, weekParam string) ([]model.Schedule, error)
	Update(ctx context.Context, userID, id string, req UpdateScheduleReq) (*model.Schedule, error)
	Delete(ctx context.Context, userID, id string) error
	Complete(ctx context.Context, userID, id string) (*model.Schedule, error)
	Skip(ctx context.Context, userID, id string) (*model.Schedule, error)
	Search(ctx context.Context, userID, query string) ([]model.Schedule, error)
}

type plannerService struct {
	repo repository.PlannerRepository
}

func NewPlannerService(repo repository.PlannerRepository) PlannerService {
	return &plannerService{repo}
}

type CreateScheduleReq struct {
	Title           string  `json:"title" binding:"required"`
	Description     string  `json:"description"`
	Date            string  `json:"date" binding:"required"`
	StartTime       string  `json:"start_time" binding:"required"`
	EndTime         string  `json:"end_time" binding:"required"`
	DurationMinutes int     `json:"duration_minutes"`
	FocusDimension  string  `json:"focus_dimension"`
	TargetID        *string `json:"target_id"`
}

type UpdateScheduleReq struct {
	Title           *string `json:"title"`
	Description     *string `json:"description"`
	Date            *string `json:"date"`
	StartTime       *string `json:"start_time"`
	EndTime         *string `json:"end_time"`
	DurationMinutes *int    `json:"duration_minutes"`
	FocusDimension  *string `json:"focus_dimension"`
	Status          *string `json:"status"`
	TargetID        *string `json:"target_id"`
}

func (s *plannerService) Create(ctx context.Context, userID string, req CreateScheduleReq) (*model.Schedule, error) {
	if req.FocusDimension == "" {
		req.FocusDimension = "General"
	}
	if req.DurationMinutes == 0 {
		req.DurationMinutes = calcDuration(req.StartTime, req.EndTime)
	}
	sched := &model.Schedule{
		ScheduleID:      uuid.New().String(),
		UserID:          userID,
		Title:           req.Title,
		Description:     req.Description,
		Date:            req.Date,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		DurationMinutes: req.DurationMinutes,
		FocusDimension:  req.FocusDimension,
		Status:          "planned",
		TargetID:        req.TargetID,
	}
	err := s.repo.Create(ctx, sched)
	return sched, err
}

func (s *plannerService) GetByID(ctx context.Context, userID, id string) (*model.Schedule, error) {
	return s.repo.GetByID(ctx, id, userID)
}

func (s *plannerService) GetByWeek(ctx context.Context, userID, weekParam string) ([]model.Schedule, error) {
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

func (s *plannerService) Update(ctx context.Context, userID, id string, req UpdateScheduleReq) (*model.Schedule, error) {
	sched, err := s.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if req.Title != nil {
		sched.Title = *req.Title
	}
	if req.Description != nil {
		sched.Description = *req.Description
	}
	if req.Date != nil {
		sched.Date = *req.Date
	}
	if req.StartTime != nil {
		sched.StartTime = *req.StartTime
	}
	if req.EndTime != nil {
		sched.EndTime = *req.EndTime
	}
	if req.DurationMinutes != nil {
		sched.DurationMinutes = *req.DurationMinutes
	}
	if req.FocusDimension != nil {
		sched.FocusDimension = *req.FocusDimension
	}
	if req.Status != nil {
		sched.Status = *req.Status
	}
	if req.TargetID != nil {
		sched.TargetID = req.TargetID
	}
	if sched.DurationMinutes == 0 {
		sched.DurationMinutes = calcDuration(sched.StartTime, sched.EndTime)
	}
	err = s.repo.Update(ctx, sched)
	return sched, err
}

func (s *plannerService) Delete(ctx context.Context, userID, id string) error {
	return s.repo.Delete(ctx, id, userID)
}

func (s *plannerService) Complete(ctx context.Context, userID, id string) (*model.Schedule, error) {
	sched, err := s.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	sched.Status = "completed"
	err = s.repo.Update(ctx, sched)
	return sched, err
}

func (s *plannerService) Skip(ctx context.Context, userID, id string) (*model.Schedule, error) {
	sched, err := s.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	sched.Status = "skipped"
	err = s.repo.Update(ctx, sched)
	return sched, err
}

func (s *plannerService) Search(ctx context.Context, userID, query string) ([]model.Schedule, error) {
	return s.repo.SearchByUser(ctx, userID, query)
}

func calcDuration(startTime, endTime string) int {
	st, err1 := time.Parse("15:04", startTime)
	et, err2 := time.Parse("15:04", endTime)
	if err1 != nil || err2 != nil {
		return 60
	}
	d := int(et.Sub(st).Minutes())
	if d <= 0 {
		d = 60
	}
	return d
}
