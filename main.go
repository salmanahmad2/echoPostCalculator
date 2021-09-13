package main

import (
	"github.com/labstack/echo/v4"
	//_ "github.com/salmanahmad2/echoPostCalculator"

	// _ "github.com/swaggo/echo-swagger/example/docs"
	_ "github.com/salmanahmad2/echoPostCalculator/docs"
	"github.com/salmanahmad2/echoPostCalculator/routes"

	_ "github.com/salmanahmad2/echoPostCalculator/routes"
)

// @title Calculator
// @version 2.0
// @description Different operations of calculator.
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @host localhost:1323
// @BasePath /
func main() {
	e := echo.New()
	routes.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
