package handlers

import (
	"exchange/internal/logic"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandlerWelcome(c echo.Context) error {

	err := c.String(http.StatusOK, logic.Welcome())

	if err != nil {
		return err
	}

	return nil
}
