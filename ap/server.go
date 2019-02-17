package main

import (
	"github.com/lunarxlark/echo-learning/ap/db"
	"github.com/lunarxlark/echo-learning/ap/logger"
	"github.com/lunarxlark/echo-learning/ap/route"

	"github.com/labstack/echo/middleware"
)

func main() {

	db.Open()
	e := route.Init()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           logger.LtsvLogFormat(),
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           logger.LogFile("./log/access.log"),
	}))
	e.Logger.Fatal(e.Start(":1323"))
}
