//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/service"
)

func InitializeUserServiceHandler() (*handler.Handler, error) {
	wire.Build(
		repository.NewDB,
		repository.NewRepository,
		service.NewService,
		handler.NewHandler,
	)
	return &handler.Handler{}, nil
}
