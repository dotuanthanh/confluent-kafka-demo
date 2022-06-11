package repository

import (
	"kafka-producer/internal"
	"kafka-producer/internal/ticker"
)

type tickerRepository struct {
	tickRepo ticker.Ticker
}

func NewTickerRepository(inter *internal.Internal) TickerRepository {
	return &tickerRepository{
		tickRepo: inter.Ticker,
	}
}

func (t *tickerRepository) Start(signal chan struct{}) {
	t.tickRepo.StartCronJob(signal)
}
