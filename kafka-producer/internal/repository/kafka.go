package repository

import (
	"kafka-producer/internal"
	"kafka-producer/internal/kafka"
)

type kafkaRepo struct {
	kafkaRepo kafka.KafkaProducer
}

func NewKafkaRepo(inter *internal.Internal) KafkaRepository {
	return &kafkaRepo{
		kafkaRepo: inter.Kafka,
	}
}

func (k *kafkaRepo) Producer(msg []byte) error {
	return k.kafkaRepo.PubMessage(msg)
}
