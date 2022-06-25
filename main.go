package main

import (
	"log"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/jonathanleonhardt/traffic_simulator_microservice/app/kafka"
	"github.com/joho/godotenv"
	"github.com/jonathanleonhardt/traffic_simulator_microservice/infra/kafka"
)

// executada automaticamente antes do main
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	Kafka consumer x producer externo pelo docker com go 
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}

	/* Kafka consumer x producer externo pelo docker com go 
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
	}
	*/

	/* Kafka producer x consumer externo pelo docker com go 
	producer := kafka.NewKafkaProducer()
	kafka.Publish("Hello", "readteste", producer)

	for {
		_ = 1
	}
	*/

	/*  - Teste do leitor de arquivos e esportar das localizações
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()

	fmt.Println(stringJson[0])
	*/
}
