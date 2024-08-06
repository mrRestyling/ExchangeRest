package main

import (
	"context"
	"exchange/internal/config"
	"exchange/internal/handlers"
	"exchange/internal/logic"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.SetLogLevel(config.LevelLogs())

	defer config.CloseLogFile() // Убеждаемся, что файл журнала закрыт, когда функция main завершается

	fmt.Println()
	fmt.Println("...Сервер запущен!")

	srv := logic.New()
	hand := handlers.New(*srv)

	s := hand.SetRouts()

	go s.Start(config.Host() + config.Port())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Останавливаем сервер")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Println("Ошибка при остановке сервера")
		log.Fatal(err)
	}
	log.Println("->...Сервер остановлен...<-")
}
