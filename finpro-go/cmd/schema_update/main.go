package main

import (
	"finpro/config"
	"log"
)

func main() {
	config.ConnectDatabase()
	
	query := "ALTER TABLE result_summary ADD COLUMN planning_score INT, ADD COLUMN time_management_score INT, ADD COLUMN cognitive_score INT, ADD COLUMN reflection_score INT"
	
	err := config.DB.Exec(query).Error
	if err != nil {
		log.Fatalf("Failed to execute ALTER TABLE: %v", err)
	}
	
	log.Println("SUCCESS: New score columns added to 'result_summary' table in MySQL!")
}
