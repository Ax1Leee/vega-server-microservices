// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: micro/review/review.proto

package review

import (
	fmt "fmt"
	_ "github.com/Ax1Leee/vega-server-microservices/api/micro/movie"
	_ "github.com/Ax1Leee/vega-server-microservices/api/micro/user"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "go-micro.dev/v5/client"
	server "go-micro.dev/v5/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ReviewService service

type ReviewService interface {
	GetReviewFromUser(ctx context.Context, in *GetReviewFromUserRequest, opts ...client.CallOption) (*GetReviewFromUserResponse, error)
	GetReviewToMovie(ctx context.Context, in *GetReviewToMovieRequest, opts ...client.CallOption) (*GetReviewToMovieResponse, error)
	GetReview(ctx context.Context, in *GetReviewRequest, opts ...client.CallOption) (*GetReviewResponse, error)
	SetReview(ctx context.Context, in *SetReviewRequest, opts ...client.CallOption) (*SetReviewResponse, error)
}

type reviewService struct {
	c    client.Client
	name string
}

func NewReviewService(name string, c client.Client) ReviewService {
	return &reviewService{
		c:    c,
		name: name,
	}
}

func (c *reviewService) GetReviewFromUser(ctx context.Context, in *GetReviewFromUserRequest, opts ...client.CallOption) (*GetReviewFromUserResponse, error) {
	req := c.c.NewRequest(c.name, "ReviewService.GetReviewFromUser", in)
	out := new(GetReviewFromUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewService) GetReviewToMovie(ctx context.Context, in *GetReviewToMovieRequest, opts ...client.CallOption) (*GetReviewToMovieResponse, error) {
	req := c.c.NewRequest(c.name, "ReviewService.GetReviewToMovie", in)
	out := new(GetReviewToMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewService) GetReview(ctx context.Context, in *GetReviewRequest, opts ...client.CallOption) (*GetReviewResponse, error) {
	req := c.c.NewRequest(c.name, "ReviewService.GetReview", in)
	out := new(GetReviewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewService) SetReview(ctx context.Context, in *SetReviewRequest, opts ...client.CallOption) (*SetReviewResponse, error) {
	req := c.c.NewRequest(c.name, "ReviewService.SetReview", in)
	out := new(SetReviewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReviewService service

type ReviewServiceHandler interface {
	GetReviewFromUser(context.Context, *GetReviewFromUserRequest, *GetReviewFromUserResponse) error
	GetReviewToMovie(context.Context, *GetReviewToMovieRequest, *GetReviewToMovieResponse) error
	GetReview(context.Context, *GetReviewRequest, *GetReviewResponse) error
	SetReview(context.Context, *SetReviewRequest, *SetReviewResponse) error
}

func RegisterReviewServiceHandler(s server.Server, hdlr ReviewServiceHandler, opts ...server.HandlerOption) error {
	type reviewService interface {
		GetReviewFromUser(ctx context.Context, in *GetReviewFromUserRequest, out *GetReviewFromUserResponse) error
		GetReviewToMovie(ctx context.Context, in *GetReviewToMovieRequest, out *GetReviewToMovieResponse) error
		GetReview(ctx context.Context, in *GetReviewRequest, out *GetReviewResponse) error
		SetReview(ctx context.Context, in *SetReviewRequest, out *SetReviewResponse) error
	}
	type ReviewService struct {
		reviewService
	}
	h := &reviewServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ReviewService{h}, opts...))
}

type reviewServiceHandler struct {
	ReviewServiceHandler
}

func (h *reviewServiceHandler) GetReviewFromUser(ctx context.Context, in *GetReviewFromUserRequest, out *GetReviewFromUserResponse) error {
	return h.ReviewServiceHandler.GetReviewFromUser(ctx, in, out)
}

func (h *reviewServiceHandler) GetReviewToMovie(ctx context.Context, in *GetReviewToMovieRequest, out *GetReviewToMovieResponse) error {
	return h.ReviewServiceHandler.GetReviewToMovie(ctx, in, out)
}

func (h *reviewServiceHandler) GetReview(ctx context.Context, in *GetReviewRequest, out *GetReviewResponse) error {
	return h.ReviewServiceHandler.GetReview(ctx, in, out)
}

func (h *reviewServiceHandler) SetReview(ctx context.Context, in *SetReviewRequest, out *SetReviewResponse) error {
	return h.ReviewServiceHandler.SetReview(ctx, in, out)
}
