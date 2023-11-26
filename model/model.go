package model

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type QpapersInfo struct {
	// gorm.Model

	Subject_code  string `gorm:"type:text;not null;"`
	Subject_name  string `gorm:"type:text;not null;"`
	Semester      int64  `gorm:"type:integer;not null;"`
	Exam_type     string `gorm:"type:text;not null;"`
	Exam_occasion string `gorm:"type:text;not null;"`
	Exam_year     int64  `gorm:"type:integer;not null;"`
	Branch        string `gorm:"type:text;not null;"`
	File_path     string `gorm:"type:text;not null;"`
}

type QpapersLoc struct {
	gorm.Model

	Qpapers_id uuid.UUID `gorm:"type:uuid;primary_key;"`
	File_path  string    `gorm:"type:text;not null;"`
}

func NewDB(uri string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
