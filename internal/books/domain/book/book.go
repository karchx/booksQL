package book

import (
	"errors"

	"github.com/google/uuid"
)

type Book struct {
	uuid               string
	title              string
	description        *string
	tags               []string
	rateRecommendation float32
}

func NewBook(title string, tags []string, description *string) (*Book, error) {
	if title == "" {
		// create const error `InvalidTitleError`
		return nil, errors.New("title cannot be empty")
	}

	return &Book{
		title:       title,
		tags:        tags,
		description: description,
	}, nil
}

func (b Book) UUID() uuid.UUID {
	return uuid.MustParse(b.uuid)
}

func (b Book) Title() string {
	return b.title
}

func (b Book) Description() *string {
	return b.description
}

func (b Book) Tags() []string {
	return b.tags
}

func (b Book) RateRecommendation() float32 {
	return b.rateRecommendation
}
