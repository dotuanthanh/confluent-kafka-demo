package usecase

type Receiver interface {
	StartConsume()
	WriteFile()
}
