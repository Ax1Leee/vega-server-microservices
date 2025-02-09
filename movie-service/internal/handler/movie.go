package handler

import (
	pb "github.com/Ax1Leee/vega-server-microservices/api/micro/movie"
	"movie-service/internal/service"

	"context"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetMovie(ctx context.Context, req *pb.GetMovieRequest, rsp *pb.GetMovieResponse) error {
	result, err := h.service.GetMovie(uint(req.MovieId))
	if err != nil {
		return err
	}
	rsp.Movie = result.Movie
	return nil
}
