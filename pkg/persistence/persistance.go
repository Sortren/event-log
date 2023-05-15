package persistence

import "github.com/uptrace/bun"

func WithLimit(limit int) SelectCriteria {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Limit(limit)
	}
}
func WithOffset(offset int) SelectCriteria {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Offset(offset)
	}
}
