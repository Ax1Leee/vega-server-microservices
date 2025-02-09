//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"movie-service/internal/handler"
	"movie-service/internal/repository"
	"movie-service/internal/service"
)

func InitializeMovieServiceHandler() (*handler.Handler, error) {
	wire.Build(
		repository.NewDB,
		repository.NewRepository,
		service.NewService,
		handler.NewHandler,
	)
	return &handler.Handler{}, nil
}
