package service

import (
	"context"

	"github.com/karchx/booksQL/internal/books/app"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	bookClient, closeBookClient, err := NewBookClient()
}
