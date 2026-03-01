package controller

import (
	"finpro/service"
	"finpro/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	service service.AdminService
}

func NewAdminController(s service.AdminService) *AdminController {
	return &AdminController{s}
}

func (ctrl *AdminController) GetUsers(c *gin.Context) {
	users, err := ctrl.service.GetAllUsers()
	if err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}
	utils.ResponseJSON(c, http.StatusOK, true, "Success", users)
}

func (ctrl *AdminController) GetEvaluations(c *gin.Context) {
	evaluations, err := ctrl.service.GetAllEvaluations()
	if err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}
	utils.ResponseJSON(c, http.StatusOK, true, "Success", evaluations)
}
