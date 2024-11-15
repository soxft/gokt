package rmq

import (
	"context"
)

type QueueOption func(*QueueArgs)

func WithMaxRetries(maxRetries int) QueueOption {
	return func(q *QueueArgs) {
		q.maxRetries = maxRetries
	}
}

func WithContext(ctx context.Context) QueueOption {
	return func(q *QueueArgs) {
		q.ctx = ctx
	}
}
