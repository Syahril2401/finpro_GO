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

func (ctrl *AdminController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseJSON(c, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	token, err := ctrl.service.Login(input.Email, input.Password)
	if err != nil {
		utils.ResponseJSON(c, http.StatusUnauthorized, false, err.Error(), nil)
		return
	}

	utils.ResponseJSON(c, http.StatusOK, true, "Login successful", gin.H{"token": token})
}

