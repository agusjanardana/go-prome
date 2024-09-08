package utils

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	log.Println(err.Error())
	c.JSON(http.StatusInternalServerError, echo.Map{"errors": err.Error()})
}
