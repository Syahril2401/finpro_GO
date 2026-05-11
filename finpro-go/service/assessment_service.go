package service

import (
	"context"
	"encoding/json"
	"finpro/model"
	"finpro/repository"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AssessmentService interface {
	GetQuestions() ([]model.Question, error)
	SubmitResponses(ctx context.Context, userID string, answers []model.AssessmentInput) (*model.ResultSummary, error)
	HasCompletedAssessment(userID string) (bool, error)
}

type assessmentService struct {
	repo  repository.AssessmentRepository
	aiSvc AIService
}

func NewAssessmentService(repo repository.AssessmentRepository, ai AIService) AssessmentService {
	return &assessmentService{repo: repo, aiSvc: ai}
}

func (s *assessmentService) GetQuestions() ([]model.Question, error) {
	return s.repo.GetQuestions()
}

func (s *assessmentService) HasCompletedAssessment(userID string) (bool, error) {
	return s.repo.HasCompletedAssessment(userID)
}

func (s *assessmentService) SubmitResponses(ctx context.Context, userID string, answers []model.AssessmentInput) (*model.ResultSummary, error) {
	// 1. Calculate scores per category
	categoryScores := make(map[string]int)
	var totalScore int

	for _, ans := range answers {
		q, err := s.repo.GetQuestionByID(ans.QuestionID)
		if err != nil {
			continue
		}
		if q.Category != nil {
			categoryScores[*q.Category] += ans.AnswerValue
		}
		totalScore += ans.AnswerValue
	}

	// 2. Map dimension scores
	planning := categoryScores["Planning"]
	timeManagement := categoryScores["Time Management"]
	cognitive := categoryScores["Cognitive Strategy"]
	reflection := categoryScores["Reflection"]

	// 3. Determine category (Low/Medium/High)
	categoryLabel := func(score int) string {
		if score <= 10 {
			return "Low"
		}
		if score <= 18 {
			return "Medium"
		}
		return "High"
	}

	// Helper to convert int to *int
	intPtr := func(i int) *int { return &i }

	// 4. Create ResultSummary object
	summary := &model.ResultSummary{
		ResultID:            uuid.New().String(),
		UserID:              userID,
		TotalScore:          intPtr(totalScore),
		PlanningScore:       intPtr(planning),
		TimeManagementScore: intPtr(timeManagement),
		CognitiveScore:      intPtr(cognitive),
		ReflectionScore:     intPtr(reflection),
	}

	// 5. Build structured Gemini prompt matching frontend Result.vue expectations
	geminiPrompt := fmt.Sprintf(`
User learning profile from Self-Regulated Learning (SRL) assessment:

Planning:         %s (Score: %d/25)
Time Management:  %s (Score: %d/25)
Cognitive Strategy: %s (Score: %d/25)
Reflection:       %s (Score: %d/25)
Total Score:      %d/100

Generate a JSON response for a university student study profile. 
Return ONLY valid JSON. No markdown, no extra text.

Structure:
{
  "profile_title": "string (motivational, e.g., The Strategic Visionary)",
  "deep_work_capacity": number (0-100),
  "cognitive_style": "string (e.g., Analytical & Pattern-Oriented)",
  "consistency_score": number (0-100),
  "retention_score": number (0-100),
  "strengths": [
    { "title": "string", "desc": "short description", "icon": "SVG_PATH" }
  ],
  "weaknesses": [
    { "title": "string", "desc": "short description" }
  ],
  "areas_for_growth": [
    { "title": "string", "desc": "short description" }
  ],
  "ai_strategy": {
    "title": "string (the primary recommendation name)",
    "desc": "short description (2-3 sentences)"
  },
  "recommendations": [
    { "title": "string", "desc": "short description" }
  ]
}

Items count: 2 strengths, 2 weaknesses, 2 areas_for_growth, 1 ai_strategy, 2 recommendations.
Icons: Use standard Heroicons SVG paths (starting with M).
`,
		categoryLabel(planning), planning,
		categoryLabel(timeManagement), timeManagement,
		categoryLabel(cognitive), cognitive,
		categoryLabel(reflection), reflection,
		totalScore,
	)

	// 6. Call Gemini API
	log.Printf("[AI] Analyzing for user %s...\n", userID)
	aiResponse, err := s.aiSvc.AnalyzeLearningStyle(ctx, geminiPrompt)
	
	if err != nil || !json.Valid([]byte(aiResponse)) {
		errDesc := "Unknown Error"
		if err != nil {
			errDesc = err.Error()
		}
		log.Printf("[AI] Using fallback for user %s. Reason: %s", userID, errDesc)

		// Sanitize error message for JSON
		safeErr, _ := json.Marshal(errDesc)
		
		aiResponse = fmt.Sprintf(`{
			"profile_title": "AI Offline Mode",
			"deep_work_capacity": 92,
			"cognitive_style": "Error Mode",
			"consistency_score": 64,
			"retention_score": 88,
			"strengths": [
				{ "title": "AI Connection Issue", "desc": %s, "icon": "M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" }
			],
			"weaknesses": [],
			"areas_for_growth": [
				{ "title": "API Verification", "desc": "The system could not reach Gemini. Please check your internet or API key." }
			],
			"ai_strategy": {
				"title": "System Alert",
				"desc": "Analysis is temporarily unavailable. Using standard profile."
			},
			"recommendations": []
		}`, string(safeErr))
	} else {
		log.Printf("[AI] SUCCESS: Analysis generated for user %s\n", userID)
	}

	// 7. Store AI response as category_result JSON
	summary.CategoryResult = &aiResponse

	// 8. Save AI log
	aiLog := &model.AILog{
		AILogID:     uuid.New().String(),
		UserID:      userID,
		PromptInput: geminiPrompt,
		AIOutput:    datatypes.JSON(aiResponse),
	}

	// 9. Persist everything in a transaction
	var responses []model.AssessmentResponse
	for _, ans := range answers {
		responses = append(responses, model.AssessmentResponse{
			ResponseID:  uuid.New().String(),
			UserID:      userID,
			QuestionID:  ans.QuestionID,
			AnswerValue: ans.AnswerValue,
		})
	}

	if err := s.repo.SubmitResponses(responses, summary, aiLog); err != nil {
		log.Printf("[DB] Failed to save assessment for user %s: %v", userID, err)
		return nil, err
	}

	return summary, nil
}
