package main

import (
	"finpro/config"
	"log"
)

func main() {
	config.ConnectDatabase()

	// Drop old tables that need restructuring
	for _, tbl := range []string{"subtasks", "notifications"} {
		config.DB.Exec("DROP TABLE IF EXISTS " + tbl)
	}

	// Recreate schedules with new schema
	config.DB.Exec("DROP TABLE IF EXISTS schedules")
	scheduleSQL := `
	CREATE TABLE schedules (
		schedule_id CHAR(36) PRIMARY KEY,
		user_id CHAR(36) NOT NULL,
		title VARCHAR(200) NOT NULL,
		description TEXT,
		date DATE NOT NULL,
		start_time VARCHAR(10) NOT NULL,
		end_time VARCHAR(10) NOT NULL,
		duration_minutes INT DEFAULT 0,
		focus_dimension VARCHAR(100) DEFAULT 'General',
		status VARCHAR(50) DEFAULT 'planned',
		target_id CHAR(36),
		created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
		updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
		CONSTRAINT fk_schedules_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);`
	if err := config.DB.Exec(scheduleSQL).Error; err != nil {
		log.Fatal("Failed to create schedules:", err)
	}
	log.Println("✓ schedules table created")

	// Recreate targets with new schema
	config.DB.Exec("DROP TABLE IF EXISTS targets")
	targetSQL := `
	CREATE TABLE targets (
		target_id CHAR(36) PRIMARY KEY,
		user_id CHAR(36) NOT NULL,
		title VARCHAR(200) NOT NULL,
		description TEXT,
		focus_dimension VARCHAR(100) DEFAULT 'General',
		priority VARCHAR(20) DEFAULT 'medium',
		due_date DATE,
		status VARCHAR(50) DEFAULT 'not_started',
		progress INT DEFAULT 0,
		week_number INT,
		year INT,
		created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
		updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
		CONSTRAINT fk_targets_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);`
	if err := config.DB.Exec(targetSQL).Error; err != nil {
		log.Fatal("Failed to create targets:", err)
	}
	log.Println("✓ targets table created")

	// Create subtasks
	subtaskSQL := `
	CREATE TABLE subtasks (
		subtask_id CHAR(36) PRIMARY KEY,
		target_id CHAR(36) NOT NULL,
		user_id CHAR(36) NOT NULL,
		title VARCHAR(200) NOT NULL,
		is_completed TINYINT(1) DEFAULT 0,
		completed_at DATETIME(3),
		created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
		updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
		CONSTRAINT fk_subtasks_target FOREIGN KEY (target_id) REFERENCES targets(target_id) ON DELETE CASCADE,
		CONSTRAINT fk_subtasks_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);`
	if err := config.DB.Exec(subtaskSQL).Error; err != nil {
		log.Fatal("Failed to create subtasks:", err)
	}
	log.Println("✓ subtasks table created")

	// Create notifications
	notifSQL := `
	CREATE TABLE notifications (
		notification_id CHAR(36) PRIMARY KEY,
		user_id CHAR(36) NOT NULL,
		type VARCHAR(50) NOT NULL,
		title VARCHAR(200) NOT NULL,
		message TEXT,
		is_read TINYINT(1) DEFAULT 0,
		related_id CHAR(36),
		created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
		CONSTRAINT fk_notifications_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);`
	if err := config.DB.Exec(notifSQL).Error; err != nil {
		log.Fatal("Failed to create notifications:", err)
	}
	log.Println("✓ notifications table created")

	log.Println("SUCCESS: All workspace tables migrated!")
}
