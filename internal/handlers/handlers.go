package handlers

import (
	"exchange/internal/models"
	"log"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Serv Servicer
}

type Servicer interface {
	Welcome() string
	Business(request *models.Request) [][]int
}

func New(l Servicer) *Handlers {
	return &Handlers{
		Serv: l,
	}
}

func (h Handlers) SetRouts() *echo.Echo {

	s := echo.New()
	log.Println("<-...Сервер запущен...->")
	s.HideBanner = true

	s.GET("/welcome", h.HandlerWelcome)
	s.GET("/b2b", h.HandlerBusiness)

	return s
}
