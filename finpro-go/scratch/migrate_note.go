package main

import (
	"finpro/config"
	"log"
)

func main() {
	config.ConnectDatabase()
	
	dropSQL := "DROP TABLE IF EXISTS notes;"
	if err := config.DB.Exec(dropSQL).Error; err != nil {
		log.Println("Drop table error:", err)
	}

	createSQL := `
	CREATE TABLE notes (
		note_id CHAR(36) PRIMARY KEY,
		user_id CHAR(36) NOT NULL,
		title VARCHAR(200) DEFAULT 'Untitled',
		content_json LONGTEXT,
		content_text LONGTEXT,
		focus_dimension VARCHAR(100) DEFAULT 'General',
		tags VARCHAR(255),
		mood VARCHAR(50),
		confidence_level INT,
		planner_session_id CHAR(36),
		target_id CHAR(36),
		is_pinned TINYINT(1) DEFAULT 0,
		is_archived TINYINT(1) DEFAULT 0,
		created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
		updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
		CONSTRAINT fk_notes_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
	`
	if err := config.DB.Exec(createSQL).Error; err != nil {
		log.Fatal("Create table error:", err)
	}
	
	log.Println("SUCCESS: Note table created successfully!")
}
