package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/will3g/golang-boilerplates/go-wtih-rabbitmq/internal/infra/database"
	"github.com/will3g/golang-boilerplates/go-wtih-rabbitmq/internal/usecase"
	"github.com/will3g/golang-boilerplates/go-wtih-rabbitmq/pkg/rabbitmq"

	_ "github.com/mattn/go-sqlite3"
)

// Thread 1 - T1
func main() {
	connectionDb, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer connectionDb.Close()

	repository := database.NewOrderRepository(connectionDb)
	usecase := usecase.CalculateFinalPrice{OrderRepository: repository}

	// Creating consumer - RabbitMQ
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel) // Consuming in other thread - T2 - RabbitMQ
	rabbitmqWorker(msgRabbitmqChannel, usecase) // RabbitMQ worker - RabbitMQ
}

func rabbitmqWorker(messageChannel chan amqp.Delivery, usecaseCalculateFinalPrice usecase.CalculateFinalPrice) {
	fmt.Println("WORKER - INFO: RabbitMQ worker has started!")
	for message := range messageChannel {
		var OrderInputDTO usecase.OrderInputDTO
		err := json.Unmarshal(message.Body, &OrderInputDTO)
		if err != nil {
			panic(err)
		}
		outputDto, err := usecaseCalculateFinalPrice.Execute(OrderInputDTO)
		if err != nil {
			panic(err)
		}
		message.Ack(false)
		fmt.Printf("RabbitMQ has processed order %s\n", outputDto.ID)
	}
}
