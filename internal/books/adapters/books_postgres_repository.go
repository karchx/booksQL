package adapters

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/karchx/booksQL/internal/books/domain/book"
	"gorm.io/gorm"
)

type BookModel struct {
	UUID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title              string    `gorm:"not null"`
	Description        *string
	Tags               []string `gorm:"type:text[]"`
	RateRecommendation float32  `gorm:"not null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

type BooksPostgresRepository struct {
	db *gorm.DB
}

func NewBooksPostgresRepository(db *gorm.DB) *BooksPostgresRepository {
	return &BooksPostgresRepository{db: db}
}

func (r *BooksPostgresRepository) CreateBook(ctx context.Context, b *book.Book) error {
	bookModel := r.marshalBook(b)
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(bookModel).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *BooksPostgresRepository) marshalBook(b *book.Book) *BookModel {
	bookModel := &BookModel{
		Title:              b.Title(),
		Description:        b.Description(),
		Tags:               b.Tags(),
		RateRecommendation: b.RateRecommendation(),
	}
	return bookModel
}
