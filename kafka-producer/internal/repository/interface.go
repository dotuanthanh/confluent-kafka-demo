package repository

type KafkaRepository interface {
	Producer(msg []byte) error
}

type TickerRepository interface {
	Start(signal chan struct{})
}
