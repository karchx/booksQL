package query

import (
	"context"

	"github.com/karchx/booksQL/internal/common/auth"
	"github.com/karchx/booksQL/internal/common/decorator"
)

type BooksForUser struct {
	User auth.User
}

type BooksForUserHandler decorator.QueryHandler[BooksForUser, []Book]

type bookForUserHandler struct {
	readModel BookForUserReadModel
}

func NewBookForUserHandler(readModel BookForUserReadModel) BooksForUserHandler {
	if readModel == nil {
		panic("readModel is nil")
	}
	return decorator.ApplyQueryDecorators(
		bookForUserHandler{readModel: readModel},
	)
	// return bookForUserHandler{readModel: readModel}
}

type BookForUserReadModel interface {
	FindBookForUser(ctx context.Context, userUUID string) ([]Book, error)
}

func (h bookForUserHandler) Handle(ctx context.Context, query BooksForUser) ([]Book, error) {
	return h.readModel.FindBookForUser(ctx, query.User.UUID)
}
