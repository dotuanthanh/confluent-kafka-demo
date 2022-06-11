package repository

type IKafkaRepository interface {
	Consume() *[]byte
}
