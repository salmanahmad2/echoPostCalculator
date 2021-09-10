package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	//_ "github.com/salmanahmad2/echoPostCalculator"

	echoSwagger "github.com/swaggo/echo-swagger"
	// _ "github.com/swaggo/echo-swagger/example/docs"
	_ "github.com/salmanahmad2/echoPostCalculator/docs"
)

type Numbers struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type Response struct {
	Result float64 `json:"result"`
}

// @title Calculator
// @version 2.0
// @description Different operations of calculator.
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @host localhost:1323
// @BasePath /
func main() {
	e := echo.New()

	e.POST("/calculator/add", add)
	e.POST("/calculator/subtract", sub)
	e.POST("/calculator/multiply", multiply)
	e.POST("/calculator/divide", division)
	e.POST("/calculator/modulus", modulus)
	e.POST("/calculator/square", square)
	e.POST("/calculator/power", power)
	e.POST("/calculator/sqrt", sqrt)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func add(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)

	add := number.Number1 + number.Number2
	//result:=Response{add}
	result := Response{
		add,
	}
	return c.JSON(http.StatusOK, result)
}

func sub(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	sub := number.Number1 - number.Number2
	result := Response{
		sub,
	}
	return c.JSON(http.StatusOK, result)
}

func multiply(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	multiply := number.Number1 * number.Number2
	result := Response{
		multiply,
	}
	return c.JSON(http.StatusOK, result)
}

func division(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	var division float64 = float64(number.Number1) / float64(number.Number2)
	result := Response{
		division,
	}
	return c.JSON(http.StatusOK, result)
}

func modulus(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	modulus := int(number.Number1) % int(number.Number2)
	result := Response{
		float64(modulus),
	}
	return c.JSON(http.StatusOK, result)
}

func square(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	square1 := number.Number1 * number.Number1
	result := Response{
		square1,
	}
	return c.JSON(http.StatusOK, result)
}

func power(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	power := math.Pow(number.Number1, number.Number1)
	result := Response{
		power,
	}
	return c.JSON(http.StatusOK, result)
}

func sqrt(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	sqrt1 := math.Sqrt(number.Number1)
	result := Response{
		sqrt1,
	}
	return c.JSON(http.StatusOK, result)

}
