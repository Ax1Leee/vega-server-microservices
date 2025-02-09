package main

import (
	"github.com/Baidu-AIP/golang-sdk/aip/censor"
	"github.com/micro/plugins/v5/broker/rabbitmq"
	"go-micro.dev/v5"
	"go-micro.dev/v5/broker"

	"fmt"
	"os"
)

func main() {
	// Initialize broker
	rabbitmqBroker := rabbitmq.NewBroker(broker.Addrs(fmt.Sprintf("amqp://%s:%s@%s/", os.Getenv("RABBITMQ_USERNAME"), os.Getenv("RABBITMQ_PASSWORD"), os.Getenv("RABBITMQ_ADDRESS"))))
	if err := rabbitmqBroker.Init(); err != nil {
		panic(err)
	}
	if err := rabbitmqBroker.Connect(); err != nil {
		panic(err)
	}
	// Subscribe
	_, err := rabbitmqBroker.Subscribe("censor-service", func(event broker.Event) error {
		// Parse message
		_ = event.Message().Header["id"]
		content := string(event.Message().Body)
		// 内容审核平台-文本
		client := censor.NewClient(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))
		res := client.TextCensor(content)
		fmt.Println(res)
		return nil
	})
	if err != nil {
		panic(err)
	}

	// Initialize service
	service := micro.NewService(
		micro.Broker(rabbitmqBroker),
	)
	service.Init()
	// Run censor-service
	if err = service.Run(); err != nil {
		panic(err)
	}
}
