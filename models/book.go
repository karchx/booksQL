package models

import "gorm.io/gorm"

type Book struct {
	ID            uint   `gorm:"primary_key"`
	Title         string `gorm:"not null"`
	Author        string `gorm:"not null"`
	PublishedYear *int
	gorm.Model
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}
