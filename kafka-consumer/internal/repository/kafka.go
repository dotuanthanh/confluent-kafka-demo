package repository

import "kafka-consumer/internal/kafka"

type kafkaRepository struct {
	kafkaRepo kafka.Kafka
}

func NewKafkaRepository() IKafkaRepository {
	return &kafkaRepository{}
}

func (k *kafkaRepository) Consume() *[]byte {
	return k.kafkaRepo.SubMessage()
}
