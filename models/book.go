package models

import (
	"fmt"

	"github.com/karchx/goQL/graph/model"
	"gorm.io/gorm"
)

type Book struct {
	ID            uint   `gorm:"primary_key"`
	Title         string `gorm:"not null"`
	Author        string `gorm:"not null"`
	PublishedYear *int32
	gorm.Model
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}

func (b Book) ToGraphModel() *model.Book {
	return &model.Book{
		ID:            fmt.Sprint(b.ID),
		Title:         b.Title,
		Author:        b.Author,
		PublishedYear: *b.PublishedYear,
	}
}

// ToGraphModelBooks cast books model in graph model
func ToGraphModelBooks(books []Book) []*model.Book {
	graphBooks := make([]*model.Book, 0, len(books))
	for _, b := range books {
		graphBooks = append(graphBooks, b.ToGraphModel())
	}
	return graphBooks
}
