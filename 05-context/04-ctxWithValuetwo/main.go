package main

import (
	"context"
	"fmt"

	"github.com/disharjayanth/go-explore-interface-error-cxt-arch/05-context/04-ctxWithValuetwo/session"
)

func main() {
	ctx := context.Background()
	ctx = session.AddUserID(ctx, 1)

	userID := session.GetUserID(ctx)
	fmt.Println("userID: ", *userID)
}
