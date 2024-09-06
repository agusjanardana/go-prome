package models

import "gorm.io/gorm"

type Note struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Title     string `gorm:"type:varchar(100)" json:"title"`
	Body      string `gorm:"type:text" json:"body"`
	CreatedBy string `gorm:"type:varchar(100)" json:"created_by"`
	gorm.Model
}

func (b *Note) TableName() string {
	return "notes"
}
