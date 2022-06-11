package ticker

type Ticker interface {
	StartCronJob(control chan struct{})
	Close()
}
