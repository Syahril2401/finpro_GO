package repository

import (
	"finpro/model"

	"gorm.io/gorm"
)

type AssessmentRepository interface {
	GetQuestions() ([]model.Question, error)
	SubmitResponses(responses []model.AssessmentResponse, summary *model.ResultSummary, log *model.AILog) error
	GetLastSummary(userID string) (*model.ResultSummary, error)
	GetAllEvaluations() ([]model.Evaluation, error)
}

type assessmentRepository struct {
	db *gorm.DB
}

func NewAssessmentRepository(db *gorm.DB) AssessmentRepository {
	return &assessmentRepository{db}
}

func (r *assessmentRepository) GetQuestions() ([]model.Question, error) {
	var questions []model.Question
	err := r.db.Find(&questions).Error
	return questions, err
}

func (r *assessmentRepository) SubmitResponses(responses []model.AssessmentResponse, summary *model.ResultSummary, log *model.AILog) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Save responses
		if err := tx.Create(&responses).Error; err != nil {
			return err
		}
		// Save summary
		if err := tx.Create(summary).Error; err != nil {
			return err
		}
		// Save AI log
		if err := tx.Create(log).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *assessmentRepository) GetLastSummary(userID string) (*model.ResultSummary, error) {
	var summary model.ResultSummary
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").First(&summary).Error
	return &summary, err
}

func (r *assessmentRepository) GetAllEvaluations() ([]model.Evaluation, error) {
	var evaluations []model.Evaluation
	err := r.db.Find(&evaluations).Error
	return evaluations, err
}
