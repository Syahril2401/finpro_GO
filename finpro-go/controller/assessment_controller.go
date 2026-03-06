package controller

import (
	"finpro/model"
	"finpro/service"
	"finpro/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssessmentController struct {
	service service.AssessmentService
}

func NewAssessmentController(s service.AssessmentService) *AssessmentController {
	return &AssessmentController{s}
}

func (ctrl *AssessmentController) GetQuestions(c *gin.Context) {
	questions, err := ctrl.service.GetQuestions()
	if err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}
	utils.ResponseJSON(c, http.StatusOK, true, "Success", questions)
}

func (ctrl *AssessmentController) Submit(c *gin.Context) {
	var input []model.AssessmentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseJSON(c, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	userID := c.MustGet("userID").(string)
	summary, err := ctrl.service.SubmitResponses(c.Request.Context(), userID, input)
	if err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	utils.ResponseJSON(c, http.StatusCreated, true, "Assessment submitted successfully", summary)
}
