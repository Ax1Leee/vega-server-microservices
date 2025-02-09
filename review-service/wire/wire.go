//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"review-service/internal/handler"
	"review-service/internal/repository"
	"review-service/internal/service"
)

func InitializeReviewServiceHandler() (*handler.Handler, error) {
	wire.Build(
		repository.NewDB,
		repository.NewRepository,
		service.NewService,
		handler.NewHandler,
	)
	return &handler.Handler{}, nil
}
