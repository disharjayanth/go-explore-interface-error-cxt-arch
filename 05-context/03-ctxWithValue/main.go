package main

import (
	"context"
	"fmt"
)

type key int

// These keys are often used with ctx.WithValue as key
var userKey key = 1
var ipKey key = 2
var isAdminKey key = 3
var sessionKey key = 4

func main() {
	ctx := context.WithValue(context.Background(), userKey, 1)
	// ctx := context.Background()
	userID, ok := ctx.Value(userKey).(int)
	fmt.Println(userID, ok)
}
