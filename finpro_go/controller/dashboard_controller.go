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
	userID := c.MustGet("userID").(uint)
	data, err := ctrl.service.GetDashboardData(userID)
	if err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}
	utils.ResponseJSON(c, http.StatusOK, true, "Success", data)
}
