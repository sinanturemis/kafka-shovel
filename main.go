package main

import (
	"fmt"
	"github.com/sinanturemis/kafka-shovel/configuration"
	"github.com/sinanturemis/kafka-shovel/pkg/application"
)

func main() {
	fmt.Println("Hello, world.")
	app := application.NewApplication(configuration.ReadShovelConfig())
	app.Start()
	app.Stop()
}
