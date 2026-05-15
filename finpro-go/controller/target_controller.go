package controller

import (
	"finpro/repository"
	"finpro/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TargetController struct {
	svc service.TargetService
}

func NewTargetController(svc service.TargetService) *TargetController {
	return &TargetController{svc}
}

func (ctrl *TargetController) GetTargets(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	weekParam := c.Query("week")
	targets, err := ctrl.svc.GetByWeek(c.Request.Context(), userID, weekParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	weekStart, weekEnd := repository.WeekBounds(time.Now())
	summary, _ := ctrl.svc.GetSummary(c.Request.Context(), userID, weekStart, weekEnd)
	c.JSON(http.StatusOK, gin.H{"data": targets, "summary": summary})
}

func (ctrl *TargetController) GetTarget(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	id := c.Param("id")
	target, err := ctrl.svc.GetByID(c.Request.Context(), userID, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Target not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": target})
}

func (ctrl *TargetController) CreateTarget(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	var req service.CreateTargetReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	target, err := ctrl.svc.Create(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Target created", "data": target})
}

func (ctrl *TargetController) UpdateTarget(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	id := c.Param("id")
	var req service.UpdateTargetReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	target, err := ctrl.svc.Update(c.Request.Context(), userID, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Target updated", "data": target})
}

func (ctrl *TargetController) DeleteTarget(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	id := c.Param("id")
	err := ctrl.svc.Delete(c.Request.Context(), userID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Target deleted"})
}

func (ctrl *TargetController) CreateSubtask(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	targetID := c.Param("id")
	var req service.CreateSubtaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sub, err := ctrl.svc.CreateSubtask(c.Request.Context(), userID, targetID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Subtask created", "data": sub})
}

func (ctrl *TargetController) UpdateSubtask(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	subtaskID := c.Param("subtaskId")
	var req service.UpdateSubtaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sub, err := ctrl.svc.UpdateSubtask(c.Request.Context(), userID, subtaskID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subtask updated", "data": sub})
}

func (ctrl *TargetController) DeleteSubtask(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	subtaskID := c.Param("subtaskId")
	err := ctrl.svc.DeleteSubtask(c.Request.Context(), userID, subtaskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subtask deleted"})
}

func (ctrl *TargetController) ToggleSubtask(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	subtaskID := c.Param("subtaskId")
	sub, err := ctrl.svc.ToggleSubtask(c.Request.Context(), userID, subtaskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subtask toggled", "data": sub})
}
