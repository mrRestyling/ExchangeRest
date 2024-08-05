package main

import (
	"exchange/internal/config"
	"exchange/internal/handlers"
	"exchange/internal/logic"
	"fmt"
	"log"
)

func main() {

	config.SetLogLevel(config.LevelLogs())

	fmt.Println()
	fmt.Println("...Server running!")

	// s := handlers.SetRouts()

	srv := logic.New()
	hand := handlers.New(*srv)

	s := hand.SetRouts()

	err := s.Start(config.Host() + config.Port())
	if err != nil {
		log.Fatal(err)
	}

}
