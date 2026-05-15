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
	UserID             string         `gorm:"column:user_id;type:char(36);primaryKey"`
	Name               string         `gorm:"column:name;type:varchar(150);not null"`
	Email              string         `gorm:"column:email;type:varchar(150);not null;uniqueIndex;index:idx_users_email"`
	PasswordHash       string         `gorm:"column:password_hash;type:varchar(255);not null"`
	Level              *string        `gorm:"column:level;type:varchar(50)"`
	GoogleID           *string        `gorm:"column:google_id;type:varchar(100);uniqueIndex"`
	GoogleAccessToken  *string        `gorm:"column:google_access_token;type:text"`
	GoogleRefreshToken *string        `gorm:"column:google_refresh_token;type:text"`
	CreatedAt          time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	IsActive           bool           `gorm:"column:is_active;type:tinyint(1);not null;default:1"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;index"`

	ResultSummaries     []ResultSummary      `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	AILogs              []AILog              `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Schedules           []Schedule           `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Targets             []Target             `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Notes               []Note               `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Evaluations         []Evaluation         `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
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
	QuestionID   string    `gorm:"column:question_id;type:char(36);primaryKey" json:"question_id"`
	QuestionText string    `gorm:"column:question_text;type:text;not null" json:"question_text"`
	Category     *string   `gorm:"column:category;type:varchar(80)" json:"category"`
	ScaleMin     int       `gorm:"column:scale_min;default:1" json:"scale_min"`
	ScaleMax     int       `gorm:"column:scale_max;default:5" json:"scale_max"`
	CreatedByID  *string   `gorm:"column:created_by;type:char(36)" json:"created_by_id"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	IsActive     bool      `gorm:"column:is_active;type:tinyint(1);not null;default:1" json:"is_active"`

	CreatedBy *Admin               `gorm:"foreignKey:CreatedByID;references:AdminID;constraint:OnDelete:SET NULL"`
}

// Removed ASSESSMENT_RESPONSES

// =========================
// RESULT_SUMMARY
// =========================

type ResultSummary struct {
	ResultID            string    `gorm:"column:result_id;type:char(36);primaryKey" json:"ResultID"`
	UserID              string    `gorm:"column:user_id;type:char(36);not null;index:idx_result_user" json:"UserID"`
	SessionID           string    `gorm:"column:session_id;type:char(36);not null;index:idx_result_session" json:"SessionID"`
	TotalScore          *int      `gorm:"column:total_score" json:"TotalScore"`
	PlanningScore       *int      `gorm:"column:planning_score" json:"PlanningScore"`
	TimeManagementScore *int      `gorm:"column:time_management_score" json:"TimeManagementScore"`
	CognitiveScore      *int      `gorm:"column:cognitive_score" json:"CognitiveScore"`
	ReflectionScore     *int      `gorm:"column:reflection_score" json:"ReflectionScore"`
	CategoryResult      *string   `gorm:"column:category_result;type:text" json:"CategoryResult"` // TEXT for long AI JSON
	CreatedAt           time.Time `gorm:"column:created_at;autoCreateTime" json:"CreatedAt"`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE" json:"-"`
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
// SCHEDULES (Planner Sessions)
// =========================

type Schedule struct {
	ScheduleID      string     `gorm:"column:schedule_id;type:char(36);primaryKey" json:"id"`
	UserID          string     `gorm:"column:user_id;type:char(36);not null" json:"user_id"`
	Title           string     `gorm:"column:title;type:varchar(200);not null" json:"title"`
	Description     string     `gorm:"column:description;type:text" json:"description"`
	Date            string     `gorm:"column:date;type:date;not null" json:"date"`
	StartTime       string     `gorm:"column:start_time;type:varchar(10);not null" json:"start_time"`
	EndTime         string     `gorm:"column:end_time;type:varchar(10);not null" json:"end_time"`
	DurationMinutes int        `gorm:"column:duration_minutes;default:0" json:"duration_minutes"`
	FocusDimension  string     `gorm:"column:focus_dimension;type:varchar(100);default:'General'" json:"focus_dimension"`
	Status          string     `gorm:"column:status;type:varchar(50);default:'planned'" json:"status"`
	TargetID        *string    `gorm:"column:target_id;type:char(36)" json:"target_id"`
	CreatedAt       time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// =========================
// WEEKLY TARGETS
// =========================

type Target struct {
	TargetID       string    `gorm:"column:target_id;type:char(36);primaryKey" json:"id"`
	UserID         string    `gorm:"column:user_id;type:char(36);not null" json:"user_id"`
	Title          string    `gorm:"column:title;type:varchar(200);not null" json:"title"`
	Description    string    `gorm:"column:description;type:text" json:"description"`
	FocusDimension string    `gorm:"column:focus_dimension;type:varchar(100);default:'General'" json:"focus_dimension"`
	Priority       string    `gorm:"column:priority;type:varchar(20);default:'medium'" json:"priority"`
	DueDate        *string   `gorm:"column:due_date;type:date" json:"due_date"`
	Status         string    `gorm:"column:status;type:varchar(50);default:'not_started'" json:"status"`
	Progress       int       `gorm:"column:progress;default:0" json:"progress"`
	WeekNumber     *int      `gorm:"column:week_number" json:"week_number"`
	Year           *int      `gorm:"column:year" json:"year"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	User     User      `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE" json:"-"`
	Subtasks []Subtask `gorm:"foreignKey:TargetID;references:TargetID;constraint:OnDelete:CASCADE" json:"subtasks"`
}

// =========================
// SUBTASKS
// =========================

type Subtask struct {
	SubtaskID   string     `gorm:"column:subtask_id;type:char(36);primaryKey" json:"id"`
	TargetID    string     `gorm:"column:target_id;type:char(36);not null" json:"target_id"`
	UserID      string     `gorm:"column:user_id;type:char(36);not null" json:"user_id"`
	Title       string     `gorm:"column:title;type:varchar(200);not null" json:"title"`
	IsCompleted bool       `gorm:"column:is_completed;type:tinyint(1);default:0" json:"is_completed"`
	CompletedAt *time.Time `gorm:"column:completed_at" json:"completed_at"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Target Target `gorm:"foreignKey:TargetID;references:TargetID;constraint:OnDelete:CASCADE" json:"-"`
	User   User   `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// =========================
// NOTIFICATIONS
// =========================

type Notification struct {
	NotificationID string    `gorm:"column:notification_id;type:char(36);primaryKey" json:"id"`
	UserID         string    `gorm:"column:user_id;type:char(36);not null" json:"user_id"`
	Type           string    `gorm:"column:type;type:varchar(50);not null" json:"type"`
	Title          string    `gorm:"column:title;type:varchar(200);not null" json:"title"`
	Message        string    `gorm:"column:message;type:text" json:"message"`
	IsRead         bool      `gorm:"column:is_read;type:tinyint(1);default:0" json:"is_read"`
	RelatedID      *string   `gorm:"column:related_id;type:char(36)" json:"related_id"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// =========================
// NOTES
// =========================

type Note struct {
	NoteID           string    `gorm:"column:note_id;type:char(36);primaryKey" json:"id"`
	UserID           string    `gorm:"column:user_id;type:char(36);not null" json:"user_id"`
	Title            string    `gorm:"column:title;type:varchar(200);default:'Untitled'" json:"title"`
	ContentJSON      string    `gorm:"column:content_json;type:longtext" json:"content_json"`
	ContentText      string    `gorm:"column:content_text;type:longtext" json:"content_text"`
	FocusDimension   string    `gorm:"column:focus_dimension;type:varchar(100);default:'General'" json:"focus_dimension"`
	Tags             string    `gorm:"column:tags;type:varchar(255)" json:"tags"`
	Mood             string    `gorm:"column:mood;type:varchar(50)" json:"mood"`
	ConfidenceLevel  *int      `gorm:"column:confidence_level" json:"confidence_level"`
	PlannerSessionID *string   `gorm:"column:planner_session_id;type:char(36)" json:"planner_session_id"`
	TargetID         *string   `gorm:"column:target_id;type:char(36)" json:"target_id"`
	IsPinned         bool      `gorm:"column:is_pinned;type:tinyint(1);default:0" json:"is_pinned"`
	IsArchived       bool      `gorm:"column:is_archived;type:tinyint(1);default:0" json:"is_archived"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE" json:"-"`
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

// =========================
// INPUT MODELS (Request DTOs)
// =========================

type AssessmentInput struct {
	QuestionID  string `json:"question_id"`
	AnswerValue int    `json:"answer_value"`
}
