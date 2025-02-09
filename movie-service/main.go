package main

import (
	pb "github.com/Ax1Leee/vega-server-microservices/api/micro/movie"
	"github.com/micro/plugins/v5/registry/etcd"
	"go-micro.dev/v5"
	"movie-service/wire"
)

func main() {
	// Initialize registry
	registry := etcd.NewRegistry()
	// Initialize movie-service
	service := micro.NewService(
		micro.Name("movie-service"),
		micro.Registry(registry),
		micro.Address(":8080"),
	)
	service.Init()
	// Initialize movie-service handler
	movieServiceHandler, err := wire.InitializeMovieServiceHandler()
	if err != nil {
		panic(err)
	}
	// Register movie-service handler
	if err = pb.RegisterMovieServiceHandler(service.Server(), movieServiceHandler); err != nil {
		panic(err)
	}
	// Run movie-service
	if err = service.Run(); err != nil {
		panic(err)
	}
}
