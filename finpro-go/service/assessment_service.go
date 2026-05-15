package service

import (
	"context"
	"encoding/json"
	"finpro/model"
	"finpro/repository"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type AssessmentService interface {
	GetQuestions() ([]model.Question, error)
	SubmitResponses(ctx context.Context, userID string, answers []model.AssessmentInput) (*model.ResultSummary, error)
	HasCompletedAssessment(userID string) (bool, error)
}

type assessmentService struct {
	repo       repository.AssessmentRepository
	profileSvc ProfileService
}

func NewAssessmentService(repo repository.AssessmentRepository, profile ProfileService) AssessmentService {
	return &assessmentService{repo: repo, profileSvc: profile}
}

func (s *assessmentService) GetQuestions() ([]model.Question, error) {
	return s.repo.GetQuestions()
}

func (s *assessmentService) HasCompletedAssessment(userID string) (bool, error) {
	return s.repo.HasCompletedAssessment(userID)
}

// categoryLabel converts a raw score (max 25) into L/M/H.
func categoryLabel(score int) string {
	if score <= 10 {
		return "L"
	}
	if score <= 18 {
		return "M"
	}
	return "H"
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
	planning := categoryScores["planning"]
	timeManagement := categoryScores["time_management"]
	cognitive := categoryScores["cognitive"]
	reflection := categoryScores["reflection"]

	// Helper to convert int to *int
	intPtr := func(i int) *int { return &i }

	// 3. Build combination_id (e.g., "L-M-H-L")
	combinationID := fmt.Sprintf("%s-%s-%s-%s",
		categoryLabel(planning),
		categoryLabel(timeManagement),
		categoryLabel(cognitive),
		categoryLabel(reflection),
	)

	log.Printf("[Profile] User %s scores: Planning=%d, TimeMgmt=%d, Cognitive=%d, Reflection=%d => %s",
		userID, planning, timeManagement, cognitive, reflection, combinationID)

	// 4. Lookup the matching profile from the 81 pre-defined profiles
	profile, err := s.profileSvc.GetProfile(combinationID)
	if err != nil {
		log.Printf("[Profile] WARNING: No profile found for %s, using fallback", combinationID)
		// Fallback to a generic profile
		profile = &SRLProfile{
			CombinationID:  combinationID,
			ProfileTitle:   "The Developing Learner",
			DeepWork:       50,
			CognitiveStyle: "Adaptive",
			Consistency:    50,
			Retention:      50,
		}
		profile.Strengths = json.RawMessage(`[]`)
		profile.Weaknesses = json.RawMessage(`[{"title":"Profile Not Found","desc":"Your combination was not in our database. Please contact support."}]`)
		profile.AreasForGrowth = json.RawMessage(`[]`)
		profile.AIStrategy = json.RawMessage(`{"title":"General Strategy","desc":"Focus on building consistent study habits across all dimensions."}`)
		profile.Recommendations = json.RawMessage(`[{"title":"Seek Guidance","desc":"Speak with a study counselor to identify your unique learning needs."}]`)
	} else {
		log.Printf("[Profile] SUCCESS: Matched profile '%s' for user %s", profile.ProfileTitle, userID)
	}

	// 5. Convert profile to JSON string for category_result
	profileJSON, err := json.Marshal(profile)
	if err != nil {
		log.Printf("[Profile] ERROR: Failed to marshal profile: %v", err)
		return nil, err
	}
	profileStr := string(profileJSON)

	// 6. Create ResultSummary
	summary := &model.ResultSummary{
		ResultID:            uuid.New().String(),
		UserID:              userID,
		TotalScore:          intPtr(totalScore),
		PlanningScore:       intPtr(planning),
		TimeManagementScore: intPtr(timeManagement),
		CognitiveScore:      intPtr(cognitive),
		ReflectionScore:     intPtr(reflection),
		CategoryResult:      &profileStr,
	}

	// 7. Persist to database
	if err := s.repo.SubmitResult(summary); err != nil {
		log.Printf("[DB] Failed to save assessment for user %s: %v", userID, err)
		return nil, err
	}

	log.Printf("[Profile] Assessment saved successfully for user %s", userID)
	return summary, nil
}
