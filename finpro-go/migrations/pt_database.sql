-- PT Database Schema (MySQL)

-- Disable foreign key checks to allow dropping and recreating tables easily
SET FOREIGN_KEY_CHECKS = 0;

-- =========================
-- USERS
-- =========================
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `user_id` CHAR(36) PRIMARY KEY,
    `name` VARCHAR(150) NOT NULL,
    `email` VARCHAR(150) NOT NULL UNIQUE,
    `password_hash` VARCHAR(255) NOT NULL,
    `level` VARCHAR(50),
    `profession` VARCHAR(100),
    `google_id` VARCHAR(100) UNIQUE,
    `google_access_token` TEXT,
    `google_refresh_token` TEXT,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_active` TINYINT(1) NOT NULL DEFAULT 1,
    `deleted_at` TIMESTAMP NULL,
    INDEX `idx_users_email` (`email`),
    INDEX `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB;

-- =========================
-- ADMINS
-- =========================
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins` (
    `admin_id` CHAR(36) PRIMARY KEY,
    `name` VARCHAR(150) NOT NULL,
    `email` VARCHAR(150) NOT NULL UNIQUE,
    `password_hash` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB;

-- =========================
-- QUESTIONS
-- =========================
DROP TABLE IF EXISTS `questions`;
CREATE TABLE `questions` (
    `question_id` CHAR(36) PRIMARY KEY,
    `question_text` TEXT NOT NULL,
    `category` VARCHAR(80),
    `scale_min` INT DEFAULT 1,
    `scale_max` INT DEFAULT 5,
    `created_by` CHAR(36),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `is_active` TINYINT(1) NOT NULL DEFAULT 1,
    CONSTRAINT `fk_questions_admin` FOREIGN KEY (`created_by`) REFERENCES `admins`(`admin_id`) ON DELETE SET NULL
) ENGINE=InnoDB;

-- =========================
-- ASSESSMENT_RESPONSES
-- =========================
DROP TABLE IF EXISTS `assessment_responses`;
CREATE TABLE `assessment_responses` (
    `response_id` CHAR(36) PRIMARY KEY,
    `user_id` CHAR(36) NOT NULL,
    `question_id` CHAR(36) NOT NULL,
    `answer_value` INT NOT NULL,
    `session_id` CHAR(36) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_assess_user_session` (`user_id`, `session_id`),
    INDEX `idx_assess_q` (`question_id`),
    CONSTRAINT `fk_assess_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) ON DELETE CASCADE,
    CONSTRAINT `fk_assess_question` FOREIGN KEY (`question_id`) REFERENCES `questions`(`question_id`) ON DELETE RESTRICT
) ENGINE=InnoDB;

-- =========================
-- RESULT_SUMMARY
-- =========================
DROP TABLE IF EXISTS `result_summary`;
CREATE TABLE `result_summary` (
    `result_id` CHAR(36) PRIMARY KEY,
    `user_id` CHAR(36) NOT NULL,
    `session_id` CHAR(36) NOT NULL,
    `total_score` INT,
    `category_result` VARCHAR(150),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_result_user` (`user_id`),
    INDEX `idx_result_session` (`session_id`),
    CONSTRAINT `fk_result_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB;

-- =========================
-- RECOMMENDATIONS
-- =========================
DROP TABLE IF EXISTS `recommendations`;
CREATE TABLE `recommendations` (
    `recommendation_id` CHAR(36) PRIMARY KEY,
    `category` VARCHAR(100),
    `description` TEXT,
    `created_by` CHAR(36),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `fk_recom_admin` FOREIGN KEY (`created_by`) REFERENCES `admins`(`admin_id`) ON DELETE SET NULL
) ENGINE=InnoDB;

-- =========================
-- AI_LOGS
-- =========================
DROP TABLE IF EXISTS `ai_logs`;
CREATE TABLE `ai_logs` (
    `ai_log_id` CHAR(36) PRIMARY KEY,
    `user_id` CHAR(36) NOT NULL,
    `prompt_input` TEXT NOT NULL,
    `ai_output` JSON,
    `model` VARCHAR(100),
    `token_count` INT,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_ai_user_created` (`user_id`, `created_at`),
    CONSTRAINT `fk_ai_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB;

-- =========================
-- SCHEDULES
-- =========================
DROP TABLE IF EXISTS `schedules`;
CREATE TABLE `schedules` (
    `schedule_id` CHAR(36) PRIMARY KEY,
    `user_id` CHAR(36) NOT NULL,
    `title` VARCHAR(200) NOT NULL,
    `description` TEXT,
    `start_time` DATETIME NOT NULL,
    `end_time` DATETIME,
    `status` ENUM('pending', 'done', 'cancelled') DEFAULT 'pending',
    `google_event_id` VARCHAR(200),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_schedule_user_start` (`user_id`, `start_time`),
    CONSTRAINT `fk_schedule_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB;

-- =========================
-- TARGETS
-- =========================
DROP TABLE IF EXISTS `targets`;
CREATE TABLE `targets` (
    `target_id` CHAR(36) PRIMARY KEY,
    `user_id` CHAR(36) NOT NULL,
    `title` VARCHAR(200) NOT NULL,
    `description` TEXT,
    `week_number` INT,
    `year` INT,
    `progress` INT DEFAULT 0,
    `status` ENUM('active', 'completed', 'cancelled') DEFAULT 'active',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_targets_user` (`user_id`, `year`, `week_number`),
    CONSTRAINT `fk_target_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB;

-- =========================
-- NOTES
-- =========================
DROP TABLE IF EXISTS `notes`;
CREATE TABLE `notes` (
    `note_id` CHAR(36) PRIMARY KEY,
    `user_id` CHAR(36) NOT NULL,
    `title` VARCHAR(200),
    `content` TEXT NOT NULL,
    `category` VARCHAR(80),
    `is_favorite` TINYINT(1) DEFAULT 0,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fk_note_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB;

-- =========================
-- EVALUATIONS
-- =========================
DROP TABLE IF EXISTS `evaluations`;
CREATE TABLE `evaluations` (
    `evaluation_id` CHAR(36) PRIMARY KEY,
    `user_id` CHAR(36) NOT NULL,
    `submitted_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `meta` JSON,
    INDEX `idx_evaluations_user` (`user_id`, `submitted_at`),
    CONSTRAINT `fk_eval_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB;

-- =========================
-- EVALUATION_RESPONSES
-- =========================
DROP TABLE IF EXISTS `evaluation_responses`;
CREATE TABLE `evaluation_responses` (
    `evaluation_response_id` CHAR(36) PRIMARY KEY,
    `evaluation_id` CHAR(36) NOT NULL,
    `question_text` TEXT NOT NULL,
    `answer_value` TEXT,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_eval_resp_eval` (`evaluation_id`),
    CONSTRAINT `fk_eval_resp_eval` FOREIGN KEY (`evaluation_id`) REFERENCES `evaluations`(`evaluation_id`) ON DELETE CASCADE
) ENGINE=InnoDB;

-- Enable foreign key checks
SET FOREIGN_KEY_CHECKS = 1;
