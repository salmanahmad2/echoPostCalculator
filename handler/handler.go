package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/salmanahmad2/echoPostCalculator/routes"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Handler(e *echo.Echo) {
	e.POST("/calculator/add", routes.Add)
	e.POST("/calculator/subtract", routes.Sub)
	e.POST("/calculator/multiply", routes.Multiply)
	e.POST("/calculator/divide", routes.Division)
	e.POST("/calculator/modulus", routes.Modulus)
	e.POST("/calculator/square", routes.Square)
	e.POST("/calculator/power", routes.Power)
	e.POST("/calculator/sqrt", routes.Sqrt)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
