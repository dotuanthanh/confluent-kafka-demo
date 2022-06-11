package configs

type Kafka struct {
	Host                     string `default:"127.0.0.1:8082""`
	Topic                    string `default:"sample-topic""`
	GroupId                  string `default:"group1""`
	EnableReadAtLatestOffset bool   `default:"false"`
}
