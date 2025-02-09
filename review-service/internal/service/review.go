package service

import (
	"github.com/Ax1Leee/vega-server-microservices/api/micro/movie"
	pb "github.com/Ax1Leee/vega-server-microservices/api/micro/review"
	"github.com/Ax1Leee/vega-server-microservices/api/micro/user"
	"github.com/micro/plugins/v5/broker/rabbitmq"
	"github.com/micro/plugins/v5/registry/etcd"
	"go-micro.dev/v5"
	"go-micro.dev/v5/broker"
	"gorm.io/gorm"
	"review-service/internal/model"
	"review-service/internal/repository"

	"context"
	"errors"
	"fmt"
	"os"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetReviewFromUser(id uint) (*pb.GetReviewFromUserResponse, error) {
	review, err := s.repository.QueryReviewByID(id)
	if err != nil {
		return nil, errors.New("failed to get review")
	}

	// Initialize registry
	registry := etcd.NewRegistry()
	// Initialize service
	service := micro.NewService(
		micro.Registry(registry),
	)
	service.Init()
	// Initialize client
	client := service.Client()
	// Call movie-service
	ctx := context.Background()
	req := client.NewRequest("movie-service", "MovieService.GetMovie", &movie.GetMovieRequest{
		MovieId: uint64(review.MovieID),
	})
	rsp := &movie.GetMovieResponse{}
	if err = client.Call(ctx, req, rsp); err != nil {
		return nil, errors.New("failed to get review")
	}

	return &pb.GetReviewFromUserResponse{Movie: rsp.Movie, Review: &pb.Review{Rating: review.Rating, Content: review.Content}}, nil
}

func (s *Service) GetReviewToMovie(id uint) (*pb.GetReviewToMovieResponse, error) {
	review, err := s.repository.QueryReviewByID(id)
	if err != nil {
		return nil, errors.New("failed to get review")
	}

	// Initialize registry
	registry := etcd.NewRegistry()
	// Initialize service
	service := micro.NewService(
		micro.Registry(registry),
	)
	service.Init()
	// Initialize client
	client := service.Client()
	// Call user-service
	ctx := context.Background()
	req := client.NewRequest("user-service", "UserService.GetUser", &user.GetUserRequest{
		UserId: uint64(review.UserID),
	})
	rsp := &user.GetUserResponse{}
	if err = client.Call(ctx, req, rsp); err != nil {
		return nil, errors.New("failed to get review")
	}

	return &pb.GetReviewToMovieResponse{User: rsp.User, Review: &pb.Review{Rating: review.Rating, Content: review.Content}}, nil
}

func (s *Service) GetReview(userID uint, movieID uint) (*pb.GetReviewResponse, error) {
	review, err := s.repository.QueryReviewByUserIDAndMovieID(userID, movieID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to get review")
		} else {
			return nil, errors.New("failed to get review")
		}
	}
	return &pb.GetReviewResponse{Review: &pb.Review{Rating: review.Rating, Content: review.Content}}, nil
}

func (s *Service) SetReview(userID uint, movieID uint, rating float32, content string) (*pb.SetReviewResponse, error) {
	review, err := s.repository.QueryReviewByUserIDAndMovieID(userID, movieID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to set review")
		} else {
			review = &model.Review{
				UserID:  userID,
				MovieID: movieID,
				Rating:  rating,
				Content: content,
				Status:  "pending",
			}
			if err = s.repository.Create(review); err != nil {
				return nil, errors.New("failed to set review")
			}
		}
	} else {
		review.Rating = rating
		review.Content = content
		review.Status = "pending"
		if err = s.repository.Update(review); err != nil {
			return nil, errors.New("failed to set review")
		}
	}

	// Initialize broker
	rabbitmqBroker := rabbitmq.NewBroker(broker.Addrs(fmt.Sprintf("amqp://%s:%s@%s/", os.Getenv("RABBITMQ_USERNAME"), os.Getenv("RABBITMQ_PASSWORD"), os.Getenv("RABBITMQ_ADDRESS"))))
	if err = rabbitmqBroker.Init(); err != nil {
		panic(err)
	}
	if err = rabbitmqBroker.Connect(); err != nil {
		panic(err)
	}
	// Publish
	if err = rabbitmqBroker.Publish("censor-service", &broker.Message{
		Header: map[string]string{"id": fmt.Sprintf("%d", review.ID)},
		Body:   []byte(fmt.Sprintf("%s", content)),
	}); err != nil {
		panic(err)
	}

	return &pb.SetReviewResponse{Review: &pb.Review{Rating: review.Rating, Content: review.Content}}, nil
}
