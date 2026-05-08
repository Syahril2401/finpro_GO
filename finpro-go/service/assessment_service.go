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
	HasCompletedAssessment(userID string) (bool, error)
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

// HasCompletedAssessment checks if a user already submitted at least one assessment
func (s *assessmentService) HasCompletedAssessment(userID string) (bool, error) {
	summary, err := s.repo.GetLastSummary(userID)
	if err != nil {
		// record not found = hasn't completed
		return false, nil
	}
	return summary != nil && summary.ResultID != "", nil
}

// categoryScore maps 1-indexed questions to their dimension category.
// Questions 1-5  → planning
// Questions 6-10 → time_management
// Questions 11-15 → cognitive
// Questions 16-20 → reflection
//
// Since the DB assigns UUIDs we rely on the ordering of answers
// (the frontend sends them in order Q1–Q20). If a Question's Category
// field is already populated in the DB we use that instead.
func scoreByCategory(answers []model.AssessmentInput, questions []model.Question) (planning, timeManagement, cognitive, reflection int) {
	// Build category lookup by question_id if questions list is provided
	catMap := map[string]string{}
	for _, q := range questions {
		if q.Category != nil {
			catMap[q.QuestionID] = *q.Category
		}
	}

	for i, a := range answers {
		cat, ok := catMap[a.QuestionID]
		if !ok {
			// Fallback: infer by 1-based position in the slice
			pos := i + 1
			switch {
			case pos <= 5:
				cat = "planning"
			case pos <= 10:
				cat = "time_management"
			case pos <= 15:
				cat = "cognitive"
			default:
				cat = "reflection"
			}
		}

		switch cat {
		case "planning":
			planning += a.AnswerValue
		case "time_management":
			timeManagement += a.AnswerValue
		case "cognitive":
			cognitive += a.AnswerValue
		case "reflection":
			reflection += a.AnswerValue
		}
	}
	return
}

func categoryLabel(score int) string {
	switch {
	case score <= 12:
		return "Low"
	case score <= 18:
		return "Medium"
	default:
		return "High"
	}
}

func (s *assessmentService) SubmitResponses(ctx context.Context, userID string, answers []model.AssessmentInput) (*model.ResultSummary, error) {
	sessionID := uuid.New().String()

	// 1. Fetch questions for category mapping
	questions, _ := s.repo.GetQuestions()

	// 2. Build response records
	var responses []model.AssessmentResponse
	for _, a := range answers {
		responses = append(responses, model.AssessmentResponse{
			ResponseID:  uuid.New().String(),
			UserID:      userID,
			QuestionID:  a.QuestionID,
			AnswerValue: a.AnswerValue,
			SessionID:   sessionID,
		})
	}

	// 3. Calculate per-dimension scores
	planning, timeManagement, cognitive, reflection := scoreByCategory(answers, questions)
	totalScore := planning + timeManagement + cognitive + reflection

	// 4. Determine overall category
	overallCategory := fmt.Sprintf(
		"Planning:%s|TimeManagement:%s|Cognitive:%s|Reflection:%s",
		categoryLabel(planning),
		categoryLabel(timeManagement),
		categoryLabel(cognitive),
		categoryLabel(reflection),
	)

	summary := &model.ResultSummary{
		ResultID:            uuid.New().String(),
		UserID:              userID,
		SessionID:           sessionID,
		TotalScore:          &totalScore,
		PlanningScore:       &planning,
		TimeManagementScore: &timeManagement,
		CognitiveScore:      &cognitive,
		ReflectionScore:     &reflection,
		CategoryResult:      &overallCategory,
	}

	// 5. Build structured Gemini prompt using per-dimension categories
	geminiPrompt := fmt.Sprintf(`
User learning profile from Self-Regulated Learning (SRL) assessment:

Planning:         %s (Score: %d/25)
Time Management:  %s (Score: %d/25)
Cognitive Strategy: %s (Score: %d/25)
Reflection:       %s (Score: %d/25)
Total Score:      %d/100

Based on this SRL profile, generate a JSON response with the following structure:
{
  "profile_title": "short motivational title (e.g., The Focused Achiever)",
  "learning_analysis": "2-3 sentences analyzing the student's overall learning pattern",
  "strengths": ["strength 1", "strength 2", "strength 3"],
  "weaknesses": ["area for improvement 1", "area for improvement 2"],
  "recommendations": [
    "specific actionable recommendation 1",
    "specific actionable recommendation 2",
    "specific actionable recommendation 3",
    "specific actionable recommendation 4",
    "specific actionable recommendation 5"
  ],
  "strategy_tags": ["Pomodoro", "Spaced Repetition", "etc"]
}

Make the response concise, motivational, and practical for a university student.
Return ONLY valid JSON, no markdown, no explanation.
`,
		categoryLabel(planning), planning,
		categoryLabel(timeManagement), timeManagement,
		categoryLabel(cognitive), cognitive,
		categoryLabel(reflection), reflection,
		totalScore,
	)

	// 6. Call Gemini API
	aiResponse, err := s.aiSvc.AnalyzeLearningStyle(ctx, geminiPrompt)
	if err != nil {
		// Fallback if Gemini fails
		aiResponse = fmt.Sprintf(`{
			"profile_title": "The Growing Learner",
			"learning_analysis": "You are developing your self-regulated learning skills. Focus on building consistent habits.",
			"strengths": ["Willingness to learn", "Openness to improvement"],
			"weaknesses": ["Consistency in planning", "Self-reflection habits"],
			"recommendations": [
				"Create a fixed weekly study schedule",
				"Use the Pomodoro technique for focused sessions",
				"Review your notes after each study session",
				"Break large tasks into smaller daily goals",
				"Track your progress weekly"
			],
			"strategy_tags": ["Time Blocking", "Pomodoro", "Active Recall"]
		}`)
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
	answersJSON, _ := json.Marshal(answers)
	_ = answersJSON

	if err := s.repo.SubmitResponses(responses, summary, aiLog); err != nil {
		return nil, err
	}

	return summary, nil
}
