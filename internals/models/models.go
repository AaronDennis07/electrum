package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Name string
	Code string
	// TotalSeats    int    `json:"total_seats"`
	// AvailabeSeats int    `json:"available_seats"`
	// Batch         int    `json:"batch"`
	// Semester      int    `json:"semester"`
	// Type          string `json:"type"`
}
