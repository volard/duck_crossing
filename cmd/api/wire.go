//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"main/internal/config"
	"your-project/internal/api/handlers"
	"your-project/internal/repository"
	"your-project/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

func InitializeAPI(db *sql.DB, cfg *config.Config) (*chi.Router, error) {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
		handlers.NewUserHandler,
		setupRouter,
	)
	return nil, nil
}
