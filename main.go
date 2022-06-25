package main

import (
	"log"

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

	producer := kafka.NewKafkaProducer()
	kafka.Publish("Hello", "readteste", producer)

	for {
		_ = 1
	}

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
