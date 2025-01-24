package contextkeys

import "context"

type contextKey string

const (
	UserId contextKey = "user_id"
	Role   contextKey = "role"
)

func GetValue[T any](ctx context.Context, key contextKey) (T, bool) {
	val, ok := ctx.Value(key).(T)

	return val, ok
}
