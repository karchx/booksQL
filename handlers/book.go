package handlers

import (
	"github.com/karchx/goQL/models"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{
		db: db,
	}
}

func (h *Handler) CreateBook(b *models.Book) error {
	return h.db.Create(b).Error
}

func (h *Handler) GetBooks() (*[]models.Book, error) {
	var b []models.Book
	if err := h.db.Order("id desc").Find(&b).Error; err != nil {
		return nil, err
	}
	return &b, nil
}
