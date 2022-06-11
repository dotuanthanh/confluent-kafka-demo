package kafka

type Kafka interface {
	SubMessage() *[]byte
	Close() error
}
