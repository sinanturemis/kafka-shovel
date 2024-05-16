package application

import (
	"github.com/sinanturemis/kafka-shovel/configuration"
	"github.com/sinanturemis/kafka-shovel/shovel"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type IApplication interface {
	Start()
	Stop()
}

type application struct {
	shovelManager shovel.IShovelManager
}

func NewApplication(configs []configuration.ShovelConfig) IApplication {
	shovelsToExecute := prepareShovels(configs)
	return &application{
		shovelManager: shovel.NewShovelManager(shovelsToExecute),
	}
}

func prepareShovels(configs []configuration.ShovelConfig) []shovel.IShovel {
	var shovelsToExecute []shovel.IShovel
	for _, config := range configs {
		shovelsToExecute = append(shovelsToExecute, shovel.NewShovel(config))
	}
	return shovelsToExecute
}

func (a *application) Start() {
	a.shovelManager.Start()
	go func() {
		_ = http.ListenAndServe(":8080", nil)
	}()
}

func (a *application) Stop() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	a.shovelManager.Stop()
}
