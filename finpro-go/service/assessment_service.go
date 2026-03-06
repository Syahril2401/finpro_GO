package service

import (
	"context"
	"encoding/json"
	"finpro/model"
	"finpro/repository"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AssessmentService interface {
	GetQuestions() ([]model.Question, error)
	SubmitResponses(ctx context.Context, userID string, answers []model.AssessmentInput) (*model.ResultSummary, error)
}

type assessmentService struct {
	repo  repository.AssessmentRepository
	aiSvc AIService
}

func NewAssessmentService(repo repository.AssessmentRepository, aiSvc AIService) AssessmentService {
	return &assessmentService{repo, aiSvc}
}

func (s *assessmentService) GetQuestions() ([]model.Question, error) {
	return s.repo.GetQuestions()
}

func (s *assessmentService) SubmitResponses(ctx context.Context, userID string, answers []model.AssessmentInput) (*model.ResultSummary, error) {
	var responses []model.AssessmentResponse
	totalScore := 0
	sessionID := uuid.New().String()

	for _, a := range answers {
		responses = append(responses, model.AssessmentResponse{
			ResponseID:  uuid.New().String(),
			UserID:      userID,
			QuestionID:  a.QuestionID,
			AnswerValue: a.AnswerValue,
			SessionID:   sessionID,
		})
		totalScore += a.AnswerValue
	}

	summary := &model.ResultSummary{
		ResultID:   uuid.New().String(),
		UserID:     userID,
		SessionID:  sessionID,
		TotalScore: &totalScore,
	}

	// Call Gemini for real analysis
	assessmentJSON, _ := json.Marshal(answers)
	aiResponse, err := s.aiSvc.AnalyzeLearningStyle(ctx, string(assessmentJSON))
	if err != nil {
		// Fallback content if AI fails
		aiResponse = `{"style": "Inconclusive", "recommendations": "Keep consistent with your current study habits."}`
	}

	aiLog := &model.AILog{
		AILogID:     uuid.New().String(),
		UserID:      userID,
		PromptInput: fmt.Sprintf("Analyze assessment score: %d, Data: %s", totalScore, string(assessmentJSON)),
		AIOutput:    datatypes.JSON(aiResponse),
	}

	summary.CategoryResult = &aiResponse // Store AI response in category result for now

	if err := s.repo.SubmitResponses(responses, summary, aiLog); err != nil {
		return nil, err
	}

	return summary, nil
}
