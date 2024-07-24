package main

import (
	"exchange/internal/config"
	"exchange/internal/handlers"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {

	fmt.Println()
	fmt.Println("...Server running!")

	s := echo.New()
	s.HideBanner = true
	s.GET("/welcome", handlers.HandlerWelcome)
	s.GET("/b2b", handlers.HandlerBusiness)

	err := s.Start(config.Host() + config.Port())
	if err != nil {
		log.Fatal(err)
	}

}
