package handler

import (
	pb "github.com/Ax1Leee/vega-server-microservices/api/micro/user"
	"user-service/internal/service"

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

func (h *Handler) GetUser(ctx context.Context, req *pb.GetUserRequest, rsp *pb.GetUserResponse) error {
	result, err := h.service.GetUser(uint(req.UserId))
	if err != nil {
		return err
	}
	rsp.User = result.User
	return nil
}
