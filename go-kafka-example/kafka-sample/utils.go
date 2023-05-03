package kafka_sample

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ReadConfig(configFile string) kafka.ConfigMap {

	conf := &kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9092",
		"group.id":          "test-kafka",
		"auto.offset.reset": "earliest",
	}

	return *conf

}
