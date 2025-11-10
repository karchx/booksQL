package ports

import (
	"net/http"

	"github.com/karchx/booksQL/internal/books/app"
	"github.com/karchx/booksQL/internal/books/app/query"
	"github.com/karchx/booksQL/internal/common/auth"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) *HttpServer {
	return &HttpServer{app: app}
}

func (h HttpServer) GetBooks(w http.ResponseWriter, r *http.Request) {
	user := auth.User{}
	var appBooks []query.Book

	appBooks, err := h.app.Queries.BookForUser.Handle(r.Context(), query.BooksForUser{User: user})
}
