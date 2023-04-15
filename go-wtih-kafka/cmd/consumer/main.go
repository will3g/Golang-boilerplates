package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/will3g/golang-boilerplates/go-wtih-kafka/internal/infra/database"
	"github.com/will3g/golang-boilerplates/go-wtih-kafka/internal/usecase"
	"github.com/will3g/golang-boilerplates/go-wtih-kafka/pkg/kafka"

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

	messageChannelKafka := make(chan *ckafka.Message)

	// Creating consumer - Kafka
	topics := []string{"orders"}
	servers := "host.docker.internal:9094"
	fmt.Println("CONSUMER - INFO: Kafka consumer has started!")
	go kafka.Consume(topics, servers, messageChannelKafka) // Consuming in other thread - T2 - Kafka
	kafkaWorker(messageChannelKafka, usecase)              // Consuming in other thread - T3 - Kafka
}

func kafkaWorker(messageChannel chan *ckafka.Message, usecaseCalculateFinalPrice usecase.CalculateFinalPrice) {
	fmt.Println("WORKER - INFO: Kafka worker has started!")
	for message := range messageChannel {
		var OrderInputDTO usecase.OrderInputDTO
		err := json.Unmarshal(message.Value, &OrderInputDTO)
		if err != nil {
			panic(err)
		}
		outputDto, err := usecaseCalculateFinalPrice.Execute(OrderInputDTO)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Kafka has processed order %s\n", outputDto.ID)
	}
}
