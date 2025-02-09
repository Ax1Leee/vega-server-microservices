package main

import (
	pb "github.com/Ax1Leee/vega-server-microservices/api/micro/user"
	"github.com/micro/plugins/v5/registry/etcd"
	"go-micro.dev/v5"
	"user-service/wire"
)

func main() {
	// Initialize registry
	registry := etcd.NewRegistry()
	// Initialize user-service
	service := micro.NewService(
		micro.Name("user-service"),
		micro.Registry(registry),
		micro.Address(":8080"),
	)
	service.Init()
	// Initialize user-service handler
	userServiceHandler, err := wire.InitializeUserServiceHandler()
	if err != nil {
		panic(err)
	}
	// Register user-service handler
	if err = pb.RegisterUserServiceHandler(service.Server(), userServiceHandler); err != nil {
		panic(err)
	}
	// Run user-service
	if err = service.Run(); err != nil {
		panic(err)
	}
}
