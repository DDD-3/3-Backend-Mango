package main

import (
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"time"
)

const timeFormat string = "2006-01-02 15:04:05 Monday"

func main() {

	//  just debug
	rand.Seed(time.Now().UnixNano())
	id := rand.Uint64()

	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{
			"server id":id, //  just debug
			"bool":true,
			"int":123,
			"float":1.0,
			"connect ip":ctx.RealIP(),
			"hello":"world",
			"time":time.Now().Format(timeFormat),
		})
	})
	e.Start(":8080")
}
