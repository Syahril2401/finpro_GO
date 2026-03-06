package controller

import (
	"finpro/config"
	"finpro/model"
	"finpro/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TestController struct{}

func NewTestController() *TestController {
	return &TestController{}
}

// AddData handles adding a question via POST (JSON) or GET (Query Parameters)
func (ctrl *TestController) AddData(c *gin.Context) {
	var question model.Question

	// Check if it's a POST request with JSON body
	if c.Request.Method == http.MethodPost {
		c.ShouldBindJSON(&question) // Try bind JSON (don't return error yet)
	}

	// Always check query params if field is still empty (fallback for POST with params)
	if question.QuestionText == "" {
		question.QuestionText = c.Query("text")
		category := c.Query("category")
		if category != "" {
			question.Category = &category
		}
	}

	// Basic validation
	if question.QuestionText == "" {
		utils.ResponseJSON(c, http.StatusBadRequest, false, "Question text is required. Use ?text=... in GET", nil)
		return
	}

	// Set ID
	question.QuestionID = uuid.New().String()
	question.IsActive = true

	// Save to DB
	if err := config.DB.Create(&question).Error; err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, false, "Failed to save data: "+err.Error(), nil)
		return
	}

	utils.ResponseJSON(c, http.StatusCreated, true, "Data added successfully", question)
}

// GetData handles selecting all questions
func (ctrl *TestController) GetData(c *gin.Context) {
	var questions []model.Question
	if err := config.DB.Find(&questions).Error; err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, false, "Failed to fetch data: "+err.Error(), nil)
		return
	}

	utils.ResponseJSON(c, http.StatusOK, true, "Success", questions)
}
