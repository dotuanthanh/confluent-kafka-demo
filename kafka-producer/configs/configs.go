package configs

import "github.com/jinzhu/configor"

type Configs struct {
	Kafka
	Ticker
}

//InitConfigs load all environment variables
func InitConfigs() (*Configs, error) {
	cfg := new(Configs)
	if err := configor.Load(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
