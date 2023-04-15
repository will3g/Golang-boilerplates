package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, servers string, messageChannel chan *ckafka.Message) {
	kafkaConsumer, err := ckafka.NewConsumer(&ckafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "whatsapp-gpt",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	err = kafkaConsumer.SubscribeTopics(topics, nil)
	if err != nil {
		panic(err)
	}
	for {
		message, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			messageChannel <- message
		}
	}
}
