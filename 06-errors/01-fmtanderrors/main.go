package main

import (
	"errors"
	"fmt"
)

var errorUser error = errors.New("User not found")
var errorID error = fmt.Errorf("User id %v not found", 1)

func main() {
	err := errorUser
	fmt.Println(errorID)

	if err == errorUser {
		fmt.Println("errors.New are same")
	}

	err = errorID
	if err == errorID {
		fmt.Println("fmt.Errorf are same")
	}
}
