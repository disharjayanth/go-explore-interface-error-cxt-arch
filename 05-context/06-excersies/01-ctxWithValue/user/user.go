package user

import (
	"context"
)

type key int

const userID key = 0
const userName key = 1

// AddUserID adds in a userid for given ctx
func AddUserID(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userID, id)
}

// AddUserName adds in a username for given ctx
func AddUserName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, userName, name)
}

// GetUserID gets id for given ctx
func GetUserID(ctx context.Context) *int {
	id, ok := ctx.Value(userID).(int)
	if !ok {
		return nil
	}
	return &id
}

// GetUserName gets name for given ctx
func GetUserName(ctx context.Context) *string {
	name, ok := ctx.Value(userName).(string)
	if !ok {
		return nil
	}
	return &name
}
