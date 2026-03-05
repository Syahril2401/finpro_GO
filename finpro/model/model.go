package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// =========================
// USERS
// =========================

type User struct {
	UserID            string         `gorm:"column:user_id;type:char(36);primaryKey"`
	Name              string         `gorm:"column:name;type:varchar(150);not null"`
	Email             string         `gorm:"column:email;type:varchar(150);not null;uniqueIndex;index:idx_users_email"`
	PasswordHash      string         `gorm:"column:password_hash;type:varchar(255);not null"`
	Level             *string        `gorm:"column:level;type:varchar(50)"`
	Profession        *string        `gorm:"column:profession;type:varchar(100)"`
	GoogleID          *string        `gorm:"column:google_id;type:varchar(100);uniqueIndex"`
	GoogleAccessToken *string        `gorm:"column:google_access_token;type:text"`
	GoogleRefreshToken *string       `gorm:"column:google_refresh_token;type:text"`
	CreatedAt         time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	IsActive          bool           `gorm:"column:is_active;type:tinyint(1);not null;default:1"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at;index"`

	AssessmentResponses []AssessmentResponse `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	ResultSummaries     []ResultSummary     `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	AILogs              []AILog             `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Schedules           []Schedule          `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Targets             []Target            `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Notes               []Note              `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Evaluations         []Evaluation        `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}

// =========================
// ADMINS (role dihapus; hanya admin)
// =========================

type Admin struct {
	AdminID      string    `gorm:"column:admin_id;type:char(36);primaryKey"`
	Name         string    `gorm:"column:name;type:varchar(150);not null"`
	Email        string    `gorm:"column:email;type:varchar(150);not null;uniqueIndex"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Questions       []Question       `gorm:"foreignKey:CreatedByID;references:AdminID"`
	Recommendations []Recommendation `gorm:"foreignKey:CreatedByID;references:AdminID"`
}

// =========================
// QUESTIONS
// =========================

type Question struct {
	QuestionID   string    `gorm:"column:question_id;type:char(36);primaryKey"`
	QuestionText string    `gorm:"column:question_text;type:text;not null"`
	Category     *string   `gorm:"column:category;type:varchar(80)"`
	ScaleMin     int       `gorm:"column:scale_min;default:1"`
	ScaleMax     int       `gorm:"column:scale_max;default:5"`
	CreatedByID  *string   `gorm:"column:created_by;type:char(36)"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	IsActive     bool      `gorm:"column:is_active;type:tinyint(1);not null;default:1"`

	CreatedBy *Admin `gorm:"foreignKey:CreatedByID;references:AdminID;constraint:OnDelete:SET NULL"`
	Responses []AssessmentResponse `gorm:"foreignKey:QuestionID;references:QuestionID;constraint:OnDelete:RESTRICT"`
}

// =========================
// ASSESSMENT_RESPONSES
// =========================

type AssessmentResponse struct {
	ResponseID  string    `gorm:"column:response_id;type:char(36);primaryKey"`
	UserID      string    `gorm:"column:user_id;type:char(36);not null;index:idx_assess_user_session,priority:1"`
	QuestionID  string    `gorm:"column:question_id;type:char(36);not null;index:idx_assess_q"`
	AnswerValue int       `gorm:"column:answer_value;not null"`
	SessionID   string    `gorm:"column:session_id;type:char(36);not null;index:idx_assess_user_session,priority:2"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`

	User     User     `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Question Question `gorm:"foreignKey:QuestionID;references:QuestionID;constraint:OnDelete:RESTRICT"`
}

// =========================
// RESULT_SUMMARY
// =========================

type ResultSummary struct {
	ResultID       string    `gorm:"column:result_id;type:char(36);primaryKey"`
	UserID         string    `gorm:"column:user_id;type:char(36);not null;index:idx_result_user"`
	SessionID      string    `gorm:"column:session_id;type:char(36);not null;index:idx_result_session"`
	TotalScore     *int      `gorm:"column:total_score"`
	CategoryResult *string   `gorm:"column:category_result;type:varchar(150)"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}

func (ResultSummary) TableName() string { return "result_summary" }

// =========================
// RECOMMENDATIONS
// =========================

type Recommendation struct {
	RecommendationID string    `gorm:"column:recommendation_id;type:char(36);primaryKey"`
	Category         *string   `gorm:"column:category;type:varchar(100)"`
	Description      *string   `gorm:"column:description;type:text"`
	CreatedByID      *string   `gorm:"column:created_by;type:char(36)"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`

	CreatedBy *Admin `gorm:"foreignKey:CreatedByID;references:AdminID;constraint:OnDelete:SET NULL"`
}

// =========================
// AI_LOGS
// =========================

type AILog struct {
	AILogID     string         `gorm:"column:ai_log_id;type:char(36);primaryKey"`
	UserID      string         `gorm:"column:user_id;type:char(36);not null;index:idx_ai_user_created,priority:1"`
	PromptInput string         `gorm:"column:prompt_input;type:text;not null"`
	AIOutput    datatypes.JSON `gorm:"column:ai_output;type:json"`
	Model       *string        `gorm:"column:model;type:varchar(100)"`
	TokenCount  *int           `gorm:"column:token_count"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime;index:idx_ai_user_created,priority:2"`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}

// =========================
// SCHEDULES
// =========================

type ScheduleStatus string

const (
	SchedulePending   ScheduleStatus = "pending"
	ScheduleDone      ScheduleStatus = "done"
	ScheduleCancelled ScheduleStatus = "cancelled"
)

type Schedule struct {
	ScheduleID    string         `gorm:"column:schedule_id;type:char(36);primaryKey"`
	UserID        string         `gorm:"column:user_id;type:char(36);not null;index:idx_schedule_user_start,priority:1"`
	Title         string         `gorm:"column:title;type:varchar(200);not null"`
	Description   *string        `gorm:"column:description;type:text"`
	StartTime     time.Time      `gorm:"column:start_time;not null;index:idx_schedule_user_start,priority:2"`
	EndTime       *time.Time     `gorm:"column:end_time"`
	Status        ScheduleStatus `gorm:"column:status;type:enum('pending','done','cancelled');default:'pending'"`
	GoogleEventID *string        `gorm:"column:google_event_id;type:varchar(200)"`
	CreatedAt     time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;autoUpdateTime"`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}

// =========================
// TARGETS
// =========================

type TargetStatus string

const (
	TargetActive    TargetStatus = "active"
	TargetCompleted TargetStatus = "completed"
	TargetCancelled TargetStatus = "cancelled"
)

type Target struct {
	TargetID    string       `gorm:"column:target_id;type:char(36);primaryKey"`
	UserID      string       `gorm:"column:user_id;type:char(36);not null;index:idx_targets_user,priority:1"`
	Title       string       `gorm:"column:title;type:varchar(200);not null"`
	Description *string      `gorm:"column:description;type:text"`
	WeekNumber  *int         `gorm:"column:week_number;index:idx_targets_user,priority:3"`
	Year        *int         `gorm:"column:year;index:idx_targets_user,priority:2"`
	Progress    int          `gorm:"column:progress;default:0"`
	Status      TargetStatus `gorm:"column:status;type:enum('active','completed','cancelled');default:'active'"`
	CreatedAt   time.Time    `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"column:updated_at;autoUpdateTime"`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}

// =========================
// NOTES
// =========================

type Note struct {
	NoteID     string    `gorm:"column:note_id;type:char(36);primaryKey"`
	UserID     string    `gorm:"column:user_id;type:char(36);not null"`
	Title      *string   `gorm:"column:title;type:varchar(200)"`
	Content    string    `gorm:"column:content;type:text;not null"`
	Category   *string   `gorm:"column:category;type:varchar(80)"`
	IsFavorite bool      `gorm:"column:is_favorite;type:tinyint(1);default:0"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}

// =========================
// EVALUATIONS & EVALUATION_RESPONSES
// =========================

type Evaluation struct {
	EvaluationID string         `gorm:"column:evaluation_id;type:char(36);primaryKey"`
	UserID       string         `gorm:"column:user_id;type:char(36);not null;index:idx_evaluations_user,priority:1"`
	SubmittedAt  time.Time      `gorm:"column:submitted_at;autoCreateTime;index:idx_evaluations_user,priority:2"`
	Meta         datatypes.JSON `gorm:"column:meta;type:json"`

	User      User                 `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Responses []EvaluationResponse `gorm:"foreignKey:EvaluationID;references:EvaluationID;constraint:OnDelete:CASCADE"`
}

type EvaluationResponse struct {
	EvaluationResponseID string    `gorm:"column:evaluation_response_id;type:char(36);primaryKey"`
	EvaluationID         string    `gorm:"column:evaluation_id;type:char(36);not null;index:idx_eval_resp_eval"`
	QuestionText         string    `gorm:"column:question_text;type:text;not null"`
	AnswerValue          *string   `gorm:"column:answer_value;type:text"`
	CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`

	Evaluation Evaluation `gorm:"foreignKey:EvaluationID;references:EvaluationID;constraint:OnDelete:CASCADE"`

}
