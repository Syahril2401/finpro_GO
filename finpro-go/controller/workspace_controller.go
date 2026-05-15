package controller

import (
	"context"
	"finpro/model"
	"finpro/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type WorkspaceController struct {
	plannerRepo  repository.PlannerRepository
	targetRepo   repository.TargetRepository
	noteRepo     repository.NoteRepository
	notifRepo    repository.NotificationRepository
	assessRepo   repository.AssessmentRepository
}

func NewWorkspaceController(
	plannerRepo repository.PlannerRepository,
	targetRepo repository.TargetRepository,
	noteRepo repository.NoteRepository,
	notifRepo repository.NotificationRepository,
	assessRepo repository.AssessmentRepository,
) *WorkspaceController {
	return &WorkspaceController{plannerRepo, targetRepo, noteRepo, notifRepo, assessRepo}
}

// ──── Dashboard Aggregation ────
func (ctrl *WorkspaceController) GetDashboardMetrics(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	ctx := c.Request.Context()
	now := time.Now()
	weekStart, weekEnd := repository.WeekBounds(now)

	// Focus Sessions (completed this week)
	focusSessions, _ := ctrl.plannerRepo.CountCompletedThisWeek(ctx, userID, weekStart, weekEnd)

	// Deep Work Hours
	deepWorkMin, _ := ctrl.plannerRepo.SumDurationCompletedThisWeek(ctx, userID, weekStart, weekEnd)
	deepWorkHours := float64(deepWorkMin) / 60.0

	// Task Efficiency
	totalSub, completedSub, _ := ctrl.targetRepo.CountSubtasksByUser(ctx, userID)
	var taskEfficiency float64
	if totalSub > 0 {
		taskEfficiency = float64(completedSub) / float64(totalSub) * 100
	} else {
		totalT, completedT, _ := ctrl.targetRepo.CountByUser(ctx, userID)
		if totalT > 0 {
			taskEfficiency = float64(completedT) / float64(totalT) * 100
		}
	}

	// Consistency (active days / 7)
	activeDays := ctrl.calcActiveDays(ctx, userID, weekStart, weekEnd)
	consistency := float64(activeDays) / 7.0 * 100

	// Learning Continuity (active weeks in last 4 / 4)
	activeWeeks := 0
	for i := 0; i < 4; i++ {
		ws, we := repository.WeekBounds(now.AddDate(0, 0, -7*i))
		days := ctrl.calcActiveDays(ctx, userID, ws, we)
		if days > 0 {
			activeWeeks++
		}
	}
	continuity := float64(activeWeeks) / 4.0 * 100

	// Weekly Focus Map (sessions per day this week)
	weekSessions, _ := ctrl.plannerRepo.GetByWeek(ctx, userID, weekStart, weekEnd)
	focusMap := make(map[string]int)
	for _, s := range weekSessions {
		focusMap[s.Date]++
	}

	// Recent sessions
	recentSessions, _ := ctrl.plannerRepo.GetByWeek(ctx, userID, weekStart, weekEnd)

	// Latest assessment
	var latestResult interface{}
	result, err := ctrl.assessRepo.GetLastResult(userID)
	if err == nil {
		latestResult = result
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"focus_sessions":     focusSessions,
			"deep_work_hours":    deepWorkHours,
			"task_efficiency":    taskEfficiency,
			"consistency":        consistency,
			"learning_continuity": continuity,
			"weekly_focus_map":   focusMap,
			"recent_sessions":   recentSessions,
			"latest_result":     latestResult,
		},
	})
}

func (ctrl *WorkspaceController) calcActiveDays(ctx context.Context, userID, weekStart, weekEnd string) int {
	daySet := make(map[string]bool)

	// Completed sessions
	sessionDates, _ := ctrl.plannerRepo.GetCompletedDatesByRange(ctx, userID, weekStart, weekEnd)
	for _, d := range sessionDates {
		daySet[d] = true
	}

	// Completed subtasks
	subtaskDates, _ := ctrl.targetRepo.GetCompletedDatesByRange(ctx, userID, weekStart, weekEnd)
	for _, d := range subtaskDates {
		daySet[d] = true
	}

	// Notes created
	notes, _ := ctrl.noteRepo.GetAllByUserID(ctx, userID)
	ws, _ := time.Parse("2006-01-02", weekStart)
	we, _ := time.Parse("2006-01-02", weekEnd)
	for _, n := range notes {
		if !n.CreatedAt.Before(ws) && !n.CreatedAt.After(we.Add(24*time.Hour)) {
			daySet[n.CreatedAt.Format("2006-01-02")] = true
		}
	}

	return len(daySet)
}

// ──── Search ────
func (ctrl *WorkspaceController) Search(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	query := c.Query("q")
	ctx := c.Request.Context()

	sessions, _ := ctrl.plannerRepo.SearchByUser(ctx, userID, query)
	targets, _ := ctrl.targetRepo.SearchByUser(ctx, userID, query)
	notes, _ := ctrl.noteRepo.SearchByUserID(ctx, userID, query)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"sessions": sessions,
			"targets":  targets,
			"notes":    notes,
		},
	})
}

// ──── Notifications ────
func (ctrl *WorkspaceController) GetNotifications(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	notifs, err := ctrl.notifRepo.GetByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	unread, _ := ctrl.notifRepo.CountUnread(c.Request.Context(), userID)
	c.JSON(http.StatusOK, gin.H{"data": notifs, "unread_count": unread})
}

func (ctrl *WorkspaceController) MarkNotificationRead(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	id := c.Param("id")
	ctrl.notifRepo.MarkRead(c.Request.Context(), id, userID)
	c.JSON(http.StatusOK, gin.H{"message": "Marked as read"})
}

func (ctrl *WorkspaceController) MarkAllNotificationsRead(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	ctrl.notifRepo.MarkAllRead(c.Request.Context(), userID)
	c.JSON(http.StatusOK, gin.H{"message": "All marked as read"})
}

// ──── Progress ────
func (ctrl *WorkspaceController) GetProgress(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	ctx := c.Request.Context()
	now := time.Now()
	weekStart, weekEnd := repository.WeekBounds(now)

	// Latest and previous assessment
	var latestResult, previousResult interface{}
	var assessmentTrend []model.ResultSummary

	results, err := ctrl.assessRepo.GetAllResults(userID)
	if err == nil && len(results) > 0 {
		latestResult = results[0]
		if len(results) > 1 {
			previousResult = results[1]
		}
		assessmentTrend = results
	}

	// Consistency
	activeDays := ctrl.calcActiveDays(ctx, userID, weekStart, weekEnd)
	consistency := float64(activeDays) / 7.0 * 100

	// Target completion
	totalSub, completedSub, _ := ctrl.targetRepo.CountSubtasksByUser(ctx, userID)
	var targetCompletion float64
	if totalSub > 0 {
		targetCompletion = float64(completedSub) / float64(totalSub) * 100
	} else {
		totalT, completedT, _ := ctrl.targetRepo.CountByUser(ctx, userID)
		if totalT > 0 {
			targetCompletion = float64(completedT) / float64(totalT) * 100
		}
	}

	// Notes count
	allNotes, _ := ctrl.noteRepo.GetAllByUserID(ctx, userID)
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	var reflectionsCount int
	for _, n := range allNotes {
		if !n.CreatedAt.Before(monthStart) {
			reflectionsCount++
		}
	}

	// Dimension delta
	var dimensionDelta interface{}
	if len(results) >= 2 {
		latest := results[0]
		prev := results[1]
		dimensionDelta = gin.H{
			"planning":        safeDelta(latest.PlanningScore, prev.PlanningScore),
			"time_management": safeDelta(latest.TimeManagementScore, prev.TimeManagementScore),
			"cognitive":       safeDelta(latest.CognitiveScore, prev.CognitiveScore),
			"reflection":      safeDelta(latest.ReflectionScore, prev.ReflectionScore),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"latest_result":     latestResult,
			"previous_result":   previousResult,
			"assessment_trend":  assessmentTrend,
			"consistency":       consistency,
			"target_completion": targetCompletion,
			"reflections_count": reflectionsCount,
			"dimension_delta":   dimensionDelta,
		},
	})
}

func safeDelta(a, b *int) int {
	if a == nil || b == nil {
		return 0
	}
	return *a - *b
}
