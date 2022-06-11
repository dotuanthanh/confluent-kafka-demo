package kafka

import (
	kafkaClient "github.com/confluentinc/confluent-kafka-go/kafka"
	"kafka-producer/configs"
)

type kafka struct {
	producer    *kafkaClient.Producer
	adminClient *kafkaClient.AdminClient
}

func NewKafka(configs *configs.Kafka) (KafkaProducer, error) {
	producer, err := kafkaClient.NewProducer(
		&kafkaClient.ConfigMap{
			"bootstrap.servers": configs.Host,
			"topic":             configs.Topic,
		})
	if err != nil {
		return nil, err
	}

	return &kafka{
		producer: producer,
	}, nil
}

func (k *kafka) PubMessage([]byte) error {
	//TODO implement me
	panic("implement me")
}

func (k *kafka) Stop() error {
	k.producer.Close()
	return nil
}
