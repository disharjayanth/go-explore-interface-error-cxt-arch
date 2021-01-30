package main

import (
	"errors"
	"fmt"
)

func cat() error {
	return errors.New("Cat is an error")
}

func moo() error {
	return fmt.Errorf("Moo is an error: %w", cat())
}

func bar() error {
	return fmt.Errorf("Bar is an error: %w", moo())
}

func foo() error {
	return fmt.Errorf("Foo is an error: %w", bar())
}

func main() {
	err := foo()
	fmt.Println("From main calling in foo() func:")
	fmt.Println(err)

	fmt.Println("Since error uses fmt.Errorf with %w it has unwrap method:")

	// Unwrap is not usually used , but its used in .Is and .As
	// Also unwrap in other words goes one level down
	// %w adds in an unwrap method to error
	// %v just shows error and doesnt have unwrap method

	// This unwraps (remove) error from foo
	// Since it has %w it has unwrap method
	baseErr := errors.Unwrap(err)
	fmt.Println(baseErr)

	// This unwraps (remove) error from bar
	// Since it has %w it has unwrap method
	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)

	// This unwraps (remove) error from moo
	// Since it has %w it has unwrap method
	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)

	// This unwraps (remove) error from cat => returns nil since
	// cat doesnt have unwrap method since it doesnt have %w
	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)
}
