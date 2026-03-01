package service

import (
	"context"
	"encoding/json"
	"finpro/model"
	"finpro/repository"
	"fmt"
)

type AssessmentService interface {
	GetQuestions() ([]model.Question, error)
	SubmitResponses(ctx context.Context, userID uint, answers []model.AssessmentInput) (*model.ResultSummary, error)
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

func (s *assessmentService) SubmitResponses(ctx context.Context, userID uint, answers []model.AssessmentInput) (*model.ResultSummary, error) {
	var responses []model.AssessmentResponse
	totalScore := 0

	for _, a := range answers {
		responses = append(responses, model.AssessmentResponse{
			UserID:     userID,
			QuestionID: a.QuestionID,
			Answer:     a.Answer,
			Score:      a.Score,
		})
		totalScore += a.Score
	}

	summary := &model.ResultSummary{
		UserID:     userID,
		TotalScore: totalScore,
	}
	// Call Gemini for real analysis
	assessmentJSON, _ := json.Marshal(answers)
	aiResponse, err := s.aiSvc.AnalyzeLearningStyle(ctx, string(assessmentJSON))
	if err != nil {
		// Fallback content if AI fails
		aiResponse = `{"style": "Inconclusive", "recommendations": "Keep consistent with your current study habits."}`
	}

	aiLog := &model.AILog{
		UserID:   userID,
		Prompt:   fmt.Sprintf("Analyze assessment score: %d, Data: %s", totalScore, string(assessmentJSON)),
		Response: aiResponse,
	}

	summary.Insight = aiResponse // Store AI response in summary insight for now

	if err := s.repo.SubmitResponses(responses, summary, aiLog); err != nil {
		return nil, err
	}

	return summary, nil
}
