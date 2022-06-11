package ticker

import (
	"kafka-producer/configs"
	"time"
)

type tickTik struct {
	t time.Ticker
}

func NewTicker(config *configs.Ticker) Ticker {
	//t := time.NewTicker()
	return &tickTik{}
}
func (t tickTik) StartCronJob() {
	//TODO implement me
	panic("implement me")
}

func (t tickTik) Close() {
	t.t.Stop()
}
