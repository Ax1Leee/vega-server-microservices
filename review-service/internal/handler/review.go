package handler

import (
	pb "github.com/Ax1Leee/vega-server-microservices/api/micro/review"
	"review-service/internal/service"

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

func (h *Handler) GetReviewFromUser(ctx context.Context, req *pb.GetReviewFromUserRequest, rsp *pb.GetReviewFromUserResponse) error {
	result, err := h.service.GetReviewFromUser(uint(req.ReviewId))
	if err != nil {
		return err
	}
	rsp.Movie = result.Movie
	rsp.Review = result.Review
	return nil
}

func (h *Handler) GetReviewToMovie(ctx context.Context, req *pb.GetReviewToMovieRequest, rsp *pb.GetReviewToMovieResponse) error {
	result, err := h.service.GetReviewToMovie(uint(req.ReviewId))
	if err != nil {
		return err
	}
	rsp.User = result.User
	rsp.Review = result.Review
	return nil
}

func (h *Handler) GetReview(ctx context.Context, req *pb.GetReviewRequest, rsp *pb.GetReviewResponse) error {
	result, err := h.service.GetReview(uint(req.UserId), uint(req.MovieId))
	if err != nil {
		return err
	}
	rsp.Review = result.Review
	return nil
}

func (h *Handler) SetReview(ctx context.Context, req *pb.SetReviewRequest, rsp *pb.SetReviewResponse) error {
	result, err := h.service.SetReview(uint(req.UserId), uint(req.MovieId), req.Rating, req.Content)
	if err != nil {
		return err
	}
	rsp.Review = result.Review
	return nil
}
