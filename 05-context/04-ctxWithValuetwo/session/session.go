package session

import (
	"context"
)

type key int

var userKey key = 0

// AddUserID adds userID value to userKey
func AddUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userKey, userID)
}

// GetUserID gets userID value from given ctx and returns a pointer to it
func GetUserID(ctx context.Context) *int {
	userID, ok := ctx.Value(userKey).(int) // type assertion
	if !ok {
		return nil
	}
	return &userID
}
