package controller

import (
	"finpro/service"
	"finpro/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseJSON(c, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	token, err := ctrl.authService.Register(input.Name, input.Email, input.Password)
	if err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	utils.ResponseJSON(c, http.StatusCreated, true, "Registration successful", gin.H{"token": token})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseJSON(c, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	token, err := ctrl.authService.Login(input.Email, input.Password)
	if err != nil {
		utils.ResponseJSON(c, http.StatusUnauthorized, false, err.Error(), nil)
		return
	}

	utils.ResponseJSON(c, http.StatusOK, true, "Login successful", gin.H{"token": token})
}
