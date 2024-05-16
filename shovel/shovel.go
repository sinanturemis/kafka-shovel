package shovel

import (
	"github.com/sinanturemis/kafka-shovel/configuration"
)

type IShovel interface {
	Subscribe()
	Unsubscribe()
	Resume()
	Pause()
}

type shovel struct {
	// ErrorConsumer       ErrorConsumer
	// RetryProducer       RetryProducer
	shovelConfig configuration.ShovelConfig
}

func NewShovel(config configuration.ShovelConfig) IShovel {
	return &shovel{shovelConfig: config}
}

func (s *shovel) Subscribe() {

}

func (s *shovel) Unsubscribe() {

}

func (s *shovel) Resume() {

}

func (s *shovel) Pause() {

}
