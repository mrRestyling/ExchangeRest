package handlers

import (
	"log"
	"net/http"

	"github.com/mrRestyling/ExchangeRest/internal/models"

	"github.com/labstack/echo/v4"
)

func (h Handlers) HandlerBusiness(c echo.Context) error {

	var request models.Request
	if err := c.Bind(&request); err != nil {
		log.Println("Пришел запрос в неверном формате")
		return err
	}
	log.Println("Получен запрос:", request)
	exchanges := h.Serv.Business(&request)

	response := map[string]interface{}{ // TODO вынести отдельно
		"exchanges": exchanges,
	}

	log.Println("Отдан ответ:", response)
	return c.JSON(http.StatusOK, response)
}
