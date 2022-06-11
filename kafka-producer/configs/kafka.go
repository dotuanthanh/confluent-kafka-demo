package configs

type Kafka struct {
	Host    string `default:"127.0.0.1""`
	Topic   string `default:"sample-topic""`
	GroupID string `default:"group1""`
}
