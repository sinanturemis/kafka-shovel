package shovel

import (
	"log"
)

type IShovelManager interface {
	Start()
	Stop()
}

type shovelManager struct {
	shovelItems []IShovel
}

func NewShovelManager(shovelItems []IShovel) IShovelManager {
	return &shovelManager{shovelItems: shovelItems}
}

func (s *shovelManager) Start() {
	log.Println("Starting...")
	for _, shovelItem := range s.shovelItems {
		go shovelItem.Subscribe()
	}
}

func (s *shovelManager) Stop() {
	log.Println("Stopping...")
	for _, shovelItem := range s.shovelItems {
		shovelItem.Unsubscribe()
	}
	//s.scheduler.Stop()
	log.Println("All shovels are stopped")
}
