package middlewares

import (
	"github.com/labstack/echo/v4"
)

func Log(c echo.Context, reqBody, resBody []byte) {
	// fmt.Println(c.Request().RequestURI, c.Request().Method)
	// mongoDB
}
