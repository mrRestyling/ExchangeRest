package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handlers) HandlerWelcome(c echo.Context) error {
	log.Println("Посещение /welcome")

	err := c.String(http.StatusOK, h.Serv.Welcome())
	if err != nil {
		log.Printf("Ошибка при отправке ответа: %v", err)
		return err
	}

	return nil
}
