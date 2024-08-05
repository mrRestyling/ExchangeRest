package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handlers) HandlerWelcome(c echo.Context) error {

	err := c.String(http.StatusOK, h.Serv.Welcome())

	if err != nil {
		return err
	}

	return nil
}
