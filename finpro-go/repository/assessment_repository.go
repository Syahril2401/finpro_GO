package repository

import (
	"finpro/model"

	"gorm.io/gorm"
)

type AssessmentRepository interface {
	GetQuestions() ([]model.Question, error)
	GetQuestionByID(id string) (model.Question, error)
	SubmitResponses(summary *model.ResultSummary, log *model.AILog) error
	SubmitResult(summary *model.ResultSummary) error
	GetLastSummary(userID string) (*model.ResultSummary, error)
	HasCompletedAssessment(userID string) (bool, error)
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

func (r *assessmentRepository) SubmitResponses(summary *model.ResultSummary, log *model.AILog) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
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

func (r *assessmentRepository) SubmitResult(summary *model.ResultSummary) error {
	return r.db.Create(summary).Error
}

func (r *assessmentRepository) GetLastSummary(userID string) (*model.ResultSummary, error) {
	var summary model.ResultSummary
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").First(&summary).Error
	return &summary, err
}

func (r *assessmentRepository) GetQuestionByID(id string) (model.Question, error) {
	var q model.Question
	err := r.db.Where("question_id = ?", id).First(&q).Error
	return q, err
}

func (r *assessmentRepository) HasCompletedAssessment(userID string) (bool, error) {
	var count int64
	err := r.db.Model(&model.ResultSummary{}).Where("user_id = ?", userID).Count(&count).Error
	return count > 0, err
}

func (r *assessmentRepository) GetAllEvaluations() ([]model.Evaluation, error) {
	var evaluations []model.Evaluation
	err := r.db.Find(&evaluations).Error
	return evaluations, err
}
