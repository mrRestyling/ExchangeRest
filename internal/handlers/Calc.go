package handlers

import (
	"exchange/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handlers) HandlerBusiness(c echo.Context) error {

	var request models.Request
	if err := c.Bind(&request); err != nil {
		return err
	}

	exchanges := h.Serv.Business(&request)

	response := map[string]interface{}{ // TODO вынести отдельно
		"exchanges": exchanges,
	}

	return c.JSON(http.StatusOK, response)
}
