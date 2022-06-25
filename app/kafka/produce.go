package kafka

import (
	"encoding/json"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jonathanleonhardt/traffic_simulator_microservice/app/route"
	"github.com/jonathanleonhardt/traffic_simulator_microservice/infra/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()

	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()

	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, position := range.positions {
		kafka.Publish(position, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep( time.Millisecond * 1000)
	}
}
