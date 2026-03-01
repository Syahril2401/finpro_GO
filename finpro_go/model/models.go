package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Email     string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`
	Role      string         `gorm:"type:enum('admin', 'user');default:'user'" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Question struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Text      string    `gorm:"type:text;not null" json:"text"`
	Category  string    `gorm:"type:varchar(100)" json:"category"`
	Weight    int       `gorm:"default:1" json:"weight"`
	CreatedAt time.Time `json:"created_at"`
}

type AssessmentResponse struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `gorm:"column:user_id" json:"user_id"`
	QuestionID uint      `gorm:"column:question_id" json:"question_id"`
	Answer     string    `gorm:"type:text" json:"answer"`
	Score      int       `gorm:"column:score" json:"score"`
	CreatedAt  time.Time `json:"created_at"`
}

type ResultSummary struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `gorm:"column:user_id" json:"user_id"`
	TotalScore int       `gorm:"column:total_score" json:"total_score"`
	Insight    string    `gorm:"type:text" json:"insight"`
	CreatedAt  time.Time `json:"created_at"`
}

type Schedule struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"column:user_id" json:"user_id"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `gorm:"type:varchar(50)" json:"status"`
}

type Target struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"column:user_id" json:"user_id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `gorm:"type:varchar(50)" json:"status"`
}

type AILog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"column:user_id" json:"user_id"`
	Prompt    string    `gorm:"type:text" json:"prompt"`
	Response  string    `gorm:"type:json" json:"response"`
	CreatedAt time.Time `json:"created_at"`
}

type Evaluation struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"column:user_id" json:"user_id"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Result    string    `gorm:"type:text" json:"result"`
	CreatedAt time.Time `json:"created_at"`
}

type Note struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"column:user_id" json:"user_id"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type AssessmentInput struct {
	QuestionID uint   `json:"question_id" binding:"required"`
	Answer     string `json:"answer" binding:"required"`
	Score      int    `json:"score" binding:"required"`
}
