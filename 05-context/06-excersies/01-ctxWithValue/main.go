package main

import (
	"context"
	"fmt"

	"github.com/disharjayanth/go-explore-interface-error-cxt-arch/05-context/06-excersies/01-ctxWithValue/user"
)

// This is from example branch

func main() {
	ctx := context.Background()
	ctx = user.AddUserID(ctx, 1)
	ctx = user.AddUserName(ctx, "James Bond")

	if id := (user.GetUserID(ctx)); id != nil {
		fmt.Println("User id from context: ", *id)
	} else {
		fmt.Println("User id not found")
	}

	if name := user.GetUserName(ctx); name != nil {
		fmt.Println("User name from context: ", *name)
	} else {
		fmt.Println("User name not found")
	}
}
