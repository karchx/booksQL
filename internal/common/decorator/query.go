package decorator

import "context"

// TODO: add metrics client prometheus
func ApplyQueryDecorators[H any, R any](handler QueryHandler[H, R]) QueryHandler[H, R] {
	return handler
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
