package kafka

//KafkaProducer is interface contain kafka 's producer method
type KafkaProducer interface {
	PubMessage([]byte) error
	Stop() error
}
