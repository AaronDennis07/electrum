package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name     *string `gorm:"unique" json:"name"`
	Admin    *string `json:"admin"`
	Password *string `json:"password"`
}

type Session struct {
	gorm.Model
	Name        *string `json:"name"`
	SessionType *string `json:"session_type"`
	Courses     []Course
}
type Course struct {
	gorm.Model
	Name         *string    `json:"name"`
	Code         *string    `json:"code"`
	Seats        *uint      `json:"seats"`
	DepartmentID *uint      `json:"department_id"`
	Department   Department `json:"-"`
	SessionID    *uint      `json:"session_id"`
	Session      Session    `json:"-"`
}

type Student struct {
	Usn          string         `gorm:"primaryKey" json:"usn"`
	Name         *string        `json:"name"`
	Email        *string        `gorm:"unique" json:"email"`
	Password     *string        `json:"-"`
	DepartmentID *uint          `json:"department_id"`
	Department   Department     `json:"-"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}
type Enrollment struct {
	gorm.Model
	Course1ID *uint   `json:"course1_id"`
	Course1   Course  `json:"-"`
	Course2ID *uint   `json:"course2_id"`
	Course2   Course  `json:"-"`
	StudentID *string `json:"student_id"`
	Student   Student `json:"-"`
	SessionID *uint   `json:"session_id"`
	Session   Session `json:"-"`
}
