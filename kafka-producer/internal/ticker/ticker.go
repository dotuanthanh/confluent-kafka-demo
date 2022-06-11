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
	timeTick := time.Date(config.Year, time.Month(config.Month), config.Day, config.Hours, config.Minute, 0, 0, time.Local)
	tickTik := &tickTik{
		timeTik: timeTick,
	}
	duration := tickTik.getNextDuration()
	ticker := time.NewTicker(duration)
	tickTik.tik = ticker
	return tickTik
}

func (t *tickTik) getNextDuration() time.Duration {
	if t.timeTik.Before(time.Now()) {
		t.timeTik.Add(24 * time.Hour)
	}
	return t.timeTik.Sub(time.Now())
}

func (t *tickTik) renew() {
	t.tik.Reset(t.getNextDuration())
}

func (t *tickTik) StartCronJob(control chan struct{}) {
	for {
		select {
		case <-t.tik.C:
			t.renew()
			close(control)
		}
	}
}

func (t *tickTik) Close() {
	t.tik.Stop()
}
