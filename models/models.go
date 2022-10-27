package models

import (
	"gorm.io/gorm"
	"time"
)

type Books struct {
	gorm.Model
	ID        uint           `gorm:"primary key,autoIncrement" json:"id"`
	Author    *string        `gorm:"type:varchar(255);not null" json:"author"`
	Title     *string        `json:"title"`
	Content   *string        `gorm:"type:text;not null" json:"content"`
	Publisher *string        `json:"publisher"`
	Year      uint16         `json:"year"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
