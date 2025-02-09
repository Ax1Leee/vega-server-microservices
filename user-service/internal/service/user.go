package service

import (
	pb "github.com/Ax1Leee/vega-server-microservices/api/micro/user"
	"gorm.io/gorm"
	"user-service/internal/repository"

	"errors"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetUser(id uint) (*pb.GetUserResponse, error) {
	user, err := s.repository.QueryUserByID(id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to get user")
		} else {
			return nil, errors.New("invalid id")
		}
	}

	resp := &pb.GetUserResponse{
		User: &pb.User{
			Avatar: user.Avatar,
			Name:   user.Name,
		},
	}
	return resp, nil
}
