package main

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{
			"bool":true,
			"int":123,
			"float":1.0,
			"connect ip":ctx.RealIP(),
			"hello":"world",
			"time":time.Now().Format("2006-01-02 15:04:05 Monday"),
		})
	})
	e.Start(":8080")
}
