package ticker

import (
	"kafka-producer/configs"
	"time"
)

type tickTik struct {
	tik     *time.Ticker
	timeTik time.Time
}

func NewTicker(config *configs.Ticker) Ticker {
	//t := time.NewTicker()
	systemTime := time.Now()
	timeTick := time.Date(config.Year, time.Month(config.Month), config.Day, config.Hours, config.Minute, 0, 0, time.Local)
	duration := systemTime.Sub(timeTick)
	ticker := time.NewTicker(duration)

	return &tickTik{
		tik:     ticker,
		timeTik: timeTick,
	}
}

func (t *tickTik) renew() {
	if t.timeTik.Before(time.Now()) {
		t.timeTik.Add(24 * time.Hour)
	}
	newDuration := t.timeTik.Sub(time.Now())
	t.tik.Reset(newDuration)
}

func (t *tickTik) StartCronJob(control chan struct{}) {
	for {
		select {
		case <-t.tik.C:
			t.renew()
		}
	}
}

func (t *tickTik) Close() {
	t.tik.Stop()
}
