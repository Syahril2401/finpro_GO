package controller

import (
	"finpro/service"
	"finpro/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	service service.DashboardService
}

func NewDashboardController(s service.DashboardService) *DashboardController {
	return &DashboardController{s}
}

func (ctrl *DashboardController) GetDashboard(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	data, err := ctrl.service.GetDashboardData(userID)
	if err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}
	utils.ResponseJSON(c, http.StatusOK, true, "Success", data)
}

func (ctrl *DashboardController) GetPlanner(c *gin.Context) {
	// Simple stub for now
	utils.ResponseJSON(c, http.StatusOK, true, "Success", gin.H{"schedules": []interface{}{}})
}

func (ctrl *DashboardController) GetNotes(c *gin.Context) {
	utils.ResponseJSON(c, http.StatusOK, true, "Success", gin.H{"notes": []interface{}{}})
}

func (ctrl *DashboardController) GetWeeklyTargets(c *gin.Context) {
	utils.ResponseJSON(c, http.StatusOK, true, "Success", gin.H{"targets": []interface{}{}})
}

func (ctrl *DashboardController) GetAIStrategies(c *gin.Context) {
	utils.ResponseJSON(c, http.StatusOK, true, "Success", gin.H{"strategies": []interface{}{}})
}

func (ctrl *DashboardController) GetProgress(c *gin.Context) {
	utils.ResponseJSON(c, http.StatusOK, true, "Success", gin.H{"progress": []interface{}{}})
}

func (ctrl *DashboardController) GetSettings(c *gin.Context) {
	utils.ResponseJSON(c, http.StatusOK, true, "Success", gin.H{"settings": map[string]interface{}{}})
}
