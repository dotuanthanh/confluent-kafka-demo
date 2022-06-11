package internal

import (
	"kafka-producer/configs"
	"kafka-producer/internal/kafka"
	"kafka-producer/internal/ticker"
	"log"
)

type Internal struct {
	Kafka  kafka.KafkaProducer
	Ticker ticker.Ticker
}

func NewInternal(cfgs *configs.Configs) *Internal {
	kafka, err := kafka.NewKafka(&cfgs.Kafka)
	if err != nil {
		log.Fatal("start internal error ", err.Error())
	}
	tik := ticker.NewTicker(&cfgs.Ticker)
	return &Internal{
		Kafka:  kafka,
		Ticker: tik,
	}
}
func (i *Internal) Close() {
	i.Ticker.Close()
	err := i.Kafka.Stop()
	if err != nil {
		return
	}
}
