package ticker

type Ticker interface {
	StartCronJob()
	Close()
}
