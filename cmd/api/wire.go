//+build wireinject

package main

import (
	"github.com/google/wire"
)

// InitializeApp initializes the Fiber app and its dependencies
func InitializeApp() (*fiber.App, error) {
	wire.Build(NewHelloHandler)
	return nil, nil
}
