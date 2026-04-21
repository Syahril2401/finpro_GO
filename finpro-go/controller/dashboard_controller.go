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
	utils.ResponseJSON(c, http.StatusOK, true, "Planner endpoint reached", gin.H{"message": "Mock data for Planner"})
}

func (ctrl *DashboardController) GetNotes(c *gin.Context) {
	utils.ResponseJSON(c, http.StatusOK, true, "Notes endpoint reached", gin.H{"message": "Mock data for Notes"})
}

func (ctrl *DashboardController) GetWeeklyTargets(c *gin.Context) {
	utils.ResponseJSON(c, http.StatusOK, true, "Weekly Targets endpoint reached", gin.H{"message": "Mock data for Weekly Targets"})
}

func (ctrl *DashboardController) GetAIStrategies(c *gin.Context) {
	utils.ResponseJSON(c, http.StatusOK, true, "AI Strategies endpoint reached", gin.H{"message": "Mock data for AI Strategies"})
}

func (ctrl *DashboardController) GetProgress(c *gin.Context) {
	utils.ResponseJSON(c, http.StatusOK, true, "Progress endpoint reached", gin.H{"message": "Mock data for Progress view"})
}

func (ctrl *DashboardController) GetSettings(c *gin.Context) {
	utils.ResponseJSON(c, http.StatusOK, true, "Settings endpoint reached", gin.H{"message": "Mock data for user Settings"})
}
