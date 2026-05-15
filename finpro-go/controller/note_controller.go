package controller

import (
	"finpro/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NoteController interface {
	CreateNote(c *gin.Context)
	GetNotes(c *gin.Context)
	GetNoteByID(c *gin.Context)
	UpdateNote(c *gin.Context)
	DeleteNote(c *gin.Context)
	SearchNotes(c *gin.Context)
	TogglePin(c *gin.Context)
	ToggleArchive(c *gin.Context)
	GetTemplates(c *gin.Context)
}

type noteController struct {
	svc service.NoteService
}

func NewNoteController(svc service.NoteService) NoteController {
	return &noteController{svc: svc}
}

func getNoteUserID(c *gin.Context) string {
	userID, exists := c.Get("userID")
	if !exists {
		return ""
	}
	return userID.(string)
}

func (ctrl *noteController) CreateNote(c *gin.Context) {
	userID := getNoteUserID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req service.CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := ctrl.svc.CreateNote(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Note created successfully", "data": note})
}

func (ctrl *noteController) GetNotes(c *gin.Context) {
	userID := getNoteUserID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	notes, err := ctrl.svc.GetAllNotes(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notes})
}

func (ctrl *noteController) GetNoteByID(c *gin.Context) {
	userID := getNoteUserID(c)
	noteID := c.Param("id")

	note, err := ctrl.svc.GetNote(c.Request.Context(), userID, noteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": note})
}

func (ctrl *noteController) UpdateNote(c *gin.Context) {
	userID := getNoteUserID(c)
	noteID := c.Param("id")

	var req service.UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := ctrl.svc.UpdateNote(c.Request.Context(), userID, noteID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note updated successfully", "data": note})
}

func (ctrl *noteController) DeleteNote(c *gin.Context) {
	userID := getNoteUserID(c)
	noteID := c.Param("id")

	err := ctrl.svc.DeleteNote(c.Request.Context(), userID, noteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}

func (ctrl *noteController) SearchNotes(c *gin.Context) {
	userID := getNoteUserID(c)
	query := c.Query("q")

	notes, err := ctrl.svc.SearchNotes(c.Request.Context(), userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notes})
}

func (ctrl *noteController) TogglePin(c *gin.Context) {
	userID := getNoteUserID(c)
	noteID := c.Param("id")

	note, err := ctrl.svc.TogglePin(c.Request.Context(), userID, noteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pin toggled", "data": note})
}

func (ctrl *noteController) ToggleArchive(c *gin.Context) {
	userID := getNoteUserID(c)
	noteID := c.Param("id")

	note, err := ctrl.svc.ToggleArchive(c.Request.Context(), userID, noteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Archive toggled", "data": note})
}

func (ctrl *noteController) GetTemplates(c *gin.Context) {
	templates := []map[string]interface{}{
		{
			"id": "blank",
			"title": "Blank Note",
			"content_json": "{\"blocks\": []}",
			"content_text": "",
		},
		{
			"id": "daily",
			"title": "Daily Reflection",
			"content_json": "{\"blocks\":[{\"id\":\"t1\",\"type\":\"heading1\",\"content\":\"Daily Reflection\"},{\"id\":\"t2\",\"type\":\"paragraph\",\"content\":\"What did I study today?\"},{\"id\":\"t3\",\"type\":\"paragraph\",\"content\":\"What strategy did I use?\"},{\"id\":\"t4\",\"type\":\"paragraph\",\"content\":\"What worked well?\"},{\"id\":\"t5\",\"type\":\"paragraph\",\"content\":\"What was difficult?\"},{\"id\":\"t6\",\"type\":\"paragraph\",\"content\":\"What will I improve next time?\"}]}",
			"content_text": "Daily Reflection\nWhat did I study today?\nWhat strategy did I use?\nWhat worked well?\nWhat was difficult?\nWhat will I improve next time?",
		},
		{
			"id": "session",
			"title": "Study Session Review",
			"content_json": "{\"blocks\":[{\"id\":\"s1\",\"type\":\"heading1\",\"content\":\"Study Session Review\"},{\"id\":\"s2\",\"type\":\"paragraph\",\"content\":\"Session goal:\"},{\"id\":\"s3\",\"type\":\"paragraph\",\"content\":\"What I completed:\"},{\"id\":\"s4\",\"type\":\"paragraph\",\"content\":\"What distracted me:\"},{\"id\":\"s5\",\"type\":\"paragraph\",\"content\":\"What helped me focus:\"},{\"id\":\"s6\",\"type\":\"paragraph\",\"content\":\"Next action:\"}]}",
			"content_text": "Study Session Review\nSession goal:\nWhat I completed:\nWhat distracted me:\nWhat helped me focus:\nNext action:",
		},
		{
			"id": "weekly",
			"title": "Weekly Learning Review",
			"content_json": "{\"blocks\":[{\"id\":\"w1\",\"type\":\"heading1\",\"content\":\"Weekly Learning Review\"},{\"id\":\"w2\",\"type\":\"paragraph\",\"content\":\"This week’s main target:\"},{\"id\":\"w3\",\"type\":\"paragraph\",\"content\":\"What I completed:\"},{\"id\":\"w4\",\"type\":\"paragraph\",\"content\":\"What I struggled with:\"},{\"id\":\"w5\",\"type\":\"paragraph\",\"content\":\"Most effective strategy:\"},{\"id\":\"w6\",\"type\":\"paragraph\",\"content\":\"One improvement for next week:\"}]}",
			"content_text": "Weekly Learning Review\nThis week’s main target:\nWhat I completed:\nWhat I struggled with:\nMost effective strategy:\nOne improvement for next week:",
		},
		{
			"id": "exam",
			"title": "Exam Preparation Note",
			"content_json": "{\"blocks\":[{\"id\":\"e1\",\"type\":\"heading1\",\"content\":\"Exam Preparation Note\"},{\"id\":\"e2\",\"type\":\"paragraph\",\"content\":\"Subject/topic:\"},{\"id\":\"e3\",\"type\":\"paragraph\",\"content\":\"Key concepts:\"},{\"id\":\"e4\",\"type\":\"paragraph\",\"content\":\"Weak areas:\"},{\"id\":\"e5\",\"type\":\"paragraph\",\"content\":\"Practice plan:\"},{\"id\":\"e6\",\"type\":\"paragraph\",\"content\":\"Review schedule:\"}]}",
			"content_text": "Exam Preparation Note\nSubject/topic:\nKey concepts:\nWeak areas:\nPractice plan:\nReview schedule:",
		},
		{
			"id": "cognitive",
			"title": "Cognitive Strategy Note",
			"content_json": "{\"blocks\":[{\"id\":\"c1\",\"type\":\"heading1\",\"content\":\"Cognitive Strategy Note\"},{\"id\":\"c2\",\"type\":\"paragraph\",\"content\":\"Topic learned:\"},{\"id\":\"c3\",\"type\":\"paragraph\",\"content\":\"Strategy used:\"},{\"id\":\"c4\",\"type\":\"paragraph\",\"content\":\"Recall result:\"},{\"id\":\"c5\",\"type\":\"paragraph\",\"content\":\"Mistakes found:\"},{\"id\":\"c6\",\"type\":\"paragraph\",\"content\":\"Better strategy next time:\"}]}",
			"content_text": "Cognitive Strategy Note\nTopic learned:\nStrategy used:\nRecall result:\nMistakes found:\nBetter strategy next time:",
		},
	}
	c.JSON(http.StatusOK, gin.H{"data": templates})
}
