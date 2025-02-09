package main

import (
	"github.com/Ax1Leee/vega-server-microservices/api/grpc-gateway/movie"
	"github.com/Ax1Leee/vega-server-microservices/api/grpc-gateway/review"
	"github.com/Ax1Leee/vega-server-microservices/api/grpc-gateway/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"context"
	"fmt"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Initialize http request multiplexer
	mux := runtime.NewServeMux()
	// Do not use TLS encryption
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	}
	// Register user-service
	if err := user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("dns:///%s", os.Getenv("USER_SERVICE_HANDLER_ENDPOINT")), opts); err != nil {
		panic(err)
	}
	// Register review-service
	if err := review.RegisterReviewServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("dns:///%s", os.Getenv("REVIEW_SERVICE_HANDLER_ENDPOINT")), opts); err != nil {
		panic(err)
	}
	// Register movie-service
	if err := movie.RegisterMovieServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("dns:///%s", os.Getenv("MOVIE_SERVICE_HANDLER_ENDPOINT")), opts); err != nil {
		panic(err)
	}
	// Start Listening and serving
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
