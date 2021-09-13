package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salmanahmad2/echoPostCalculator/controller"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Handler(e *echo.Echo) {
	e.POST("/calculator/add", controller.Add)
	e.POST("/calculator/subtract", controller.Sub)
	e.POST("/calculator/multiply", controller.Multiply)
	e.POST("/calculator/divide", controller.Division)
	e.POST("/calculator/modulus", controller.Modulus)
	e.POST("/calculator/square", controller.Square)
	e.POST("/calculator/power", controller.Power)
	e.POST("/calculator/sqrt", controller.Sqrt)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
