package main

import (
	"kafka-producer/configs"
	"kafka-producer/internal"
	"kafka-producer/internal/repository"
	"log"
)

func main() {
	cfgs, err := configs.InitConfigs()
	if err != nil {
		log.Fatal("Init config error... ", err.Error())
	}
	inter := internal.NewInternal(cfgs)

	tickerRepo := repository.NewTickerRepository(inter)
	kafkaRepo := repository.NewKafkaRepo(inter)

	if cfgs.EnableCronJobs {
		isCronRunning := make(chan struct{})
		go tickerRepo.Start(isCronRunning)
		<-isCronRunning
	}

	messDemo := []byte("aaa bbb cc")
	err = kafkaRepo.Producer(messDemo)
	if err != nil {
		return
	}

	defer inter.Close()
}
