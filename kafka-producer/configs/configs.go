package configs

type Configs struct {
	Kafka
	Ticker
}

func InitConfigs() *Configs {

	return &Configs{}
}
