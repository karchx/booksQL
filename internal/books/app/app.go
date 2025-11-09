package app

import "github.com/karchx/booksQL/internal/books/app/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	BookForUser query.BookForUserHandler
}
