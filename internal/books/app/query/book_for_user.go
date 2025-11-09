package query

import (
	"context"
	"github.com/karchx/booksQL/internal/common"
)

type BooksForUser struct {
	User auth.User
}

type BookForUserHandler struct {
	readModel BookForUserReadModel
}

func NewBookForUserHandler(readModel BookForUserReadModel) BookForUserHandler {
	if readModel == nil {
		panic("readModel is nil")
	}

	return BookForUserHandler{readModel: readModel}
}

type BookForUserReadModel interface {
	FindBookForUser(ctx context.Context, userUUID string) ([]Book, error)
}

// TODO: add user auth.User
func (h *BookForUserHandler) Handle(ctx context.Context, userUUID string) ([]Book, error) {
	return h.readModel.FindBookForUser(ctx, userUUID)
}
