package main

import (
	"fmt"
	pb "github.com/Ax1Leee/vega-server-microservices/api/micro/review"
	"github.com/go-co-op/gocron/v2"
	"github.com/micro/plugins/v5/registry/etcd"
	"go-micro.dev/v5"
	"review-service/wire"
)

func main() {
	// Create a scheduler
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = scheduler.Shutdown(); err != nil {
			panic(err)
		}
	}()
	// Add a job to the scheduler
	job, err := scheduler.NewJob(
		gocron.MonthlyJob(
			1,
			gocron.NewDaysOfTheMonth(1),
			gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0)),
		),
		gocron.NewTask(
			func() {},
		),
	)
	if err != nil {
		panic(err)
	}
	// Each job has a unique id
	fmt.Println(job.ID())
	// Start the scheduler
	scheduler.Start()

	// Initialize registry
	registry := etcd.NewRegistry()
	// Initialize review-service
	service := micro.NewService(
		micro.Name("review-service"),
		micro.Registry(registry),
		micro.Address(":8080"),
	)
	service.Init()
	// Initialize review-service handler
	reviewServiceHandler, err := wire.InitializeReviewServiceHandler()
	if err != nil {
		panic(err)
	}
	// Register review-service handler
	if err = pb.RegisterReviewServiceHandler(service.Server(), reviewServiceHandler); err != nil {
		panic(err)
	}
	// Run review-service
	if err = service.Run(); err != nil {
		panic(err)
	}
}
