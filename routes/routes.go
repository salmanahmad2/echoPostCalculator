package routes

import (
	//"crypto/subtle"

	//	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"

	"github.com/salmanahmad2/echoPostCalculator/controller"
	"github.com/salmanahmad2/echoPostCalculator/middlewares"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routes(e *echo.Echo) {

	// e.Use()

	e.POST("/calculator/add", controller.Add, middlewares.Auth)
	e.POST("/calculator/subtract", controller.Sub, middlewares.Auth)
	e.POST("/calculator/multiply", controller.Multiply, middlewares.Auth)
	e.POST("/calculator/divide", controller.Division, middlewares.Auth)
	e.POST("/calculator/modulus", controller.Modulus, middlewares.Auth)
	e.POST("/calculator/square", controller.Square, middlewares.Auth)
	e.POST("/calculator/power", controller.Power, middlewares.Auth)
	e.POST("/calculator/sqrt", controller.Sqrt, middlewares.Auth)
	e.GET("/calculator/getRecord/:id", controller.GetRecord, middlewares.Auth)
	e.GET("/calculator/getAllRecords", controller.GetAllRecords, middlewares.Auth)
	e.DELETE("/calculator/deleteRecord/:id", controller.DeleteRecord, middlewares.Auth)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
