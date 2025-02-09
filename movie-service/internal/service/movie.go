package service

import (
	pb "github.com/Ax1Leee/vega-server-microservices/api/micro/movie"
	"gorm.io/gorm"
	"movie-service/internal/repository"

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

func (s *Service) GetMovie(id uint) (*pb.GetMovieResponse, error) {
	movie, err := s.repository.QueryMovieByID(id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to get movie")
		} else {
			return nil, errors.New("invalid id")
		}
	}

	resp := &pb.GetMovieResponse{
		Movie: &pb.Movie{
			Cover:        movie.Cover,
			Title:        movie.Title,
			CriticRating: movie.CriticRating,
			UserRating:   movie.UserRating,
		},
	}
	return resp, nil
}
