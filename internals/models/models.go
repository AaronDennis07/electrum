package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name     *string `gorm:"unique"`
	Admin    *string
	Password *string
}

type Session struct {
	gorm.Model
	Name        *string
	SessionType *string
}
type Course struct {
	gorm.Model
	Name         *string
	Code         *string
	DepartmentID *uint
	Department   Department
	SessionID    *uint
	Session      Session
}

type Student struct {
	Usn          string `gorm:"primaryKey"`
	Name         *string
	Email        *string `gorm:"unique"`
	Password     *string
	DepartmentID *uint
	Department   Department
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
type Enrollment struct {
	gorm.Model
	Course1ID *uint
	Course1   Course
	Course2ID *uint
	Course2   Course
	StudentID *string
	Student   Student
	SessionID *uint
	Session   Session
}
