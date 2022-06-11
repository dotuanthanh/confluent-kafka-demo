package internal

import (
	"kafka-producer/configs"
	"kafka-producer/internal/kafka"
	"log"
)

type Internal struct {
	kafka kafka.KafkaProducer
}

func NewInternal(cfgs configs.Configs) *Internal {
	kafka, err := kafka.NewKafka(&cfgs.Kafka)
	if err != nil {
		log.Fatal("start internal error ", err.Error())
	}
	return &Internal{
		kafka: kafka,
	}
}
func (i *Internal) Close() {

}
