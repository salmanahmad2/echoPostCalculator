package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Numbers struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

func main() {
	e := echo.New()

	e.POST("/calculator/add", add)
	e.POST("calculator/sub", sub)
	e.POST("/calculator/multiply", multiply)
	e.POST("/calculator/division", division)
	e.POST("/calculator/modulus", modulus)
	e.POST("/calculator/square", square)
	e.POST("/calculator/power", power)
	e.POST("/calculator/sqrt", sqrt)

	e.Logger.Fatal(e.Start(":1323"))
}
func add(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	add := number.Number1 + number.Number2
	return c.String(http.StatusOK, "Addition is:"+strconv.Itoa(add))
}

func sub(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	sub := number.Number1 - number.Number2
	return c.String(http.StatusOK, "Subtraction is:"+strconv.Itoa(sub))
}

func multiply(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	multiply := number.Number1 * number.Number2
	return c.String(http.StatusOK, "Multiplication is:"+strconv.Itoa(multiply))
}

func division(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	var division float64 = float64(number.Number1) / float64(number.Number2)
	return c.String(http.StatusOK, "Division is:"+strconv.FormatFloat(division, 'E', -1, 64))
}

func modulus(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	modulus := number.Number1 % number.Number2
	return c.String(http.StatusOK, "Modulus is:"+strconv.Itoa(modulus))
}

func square(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	square1 := number.Number1 * number.Number1
	square2 := number.Number2 * number.Number2
	return c.String(http.StatusOK, "Square of "+strconv.Itoa(number.Number1)+" is "+strconv.Itoa(square1)+"\nSquare of "+strconv.Itoa(number.Number2)+" is "+strconv.Itoa(square2))
}

func power(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	power := math.Pow(float64(number.Number1), float64(number.Number1))
	return c.String(http.StatusOK, "Number "+strconv.Itoa(number.Number1)+" raised to the power "+strconv.Itoa(number.Number2)+" is "+strconv.FormatFloat(power, 'E', -1, 64))
}

func sqrt(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	sqrt1 := math.Sqrt(float64(number.Number1))
	sqrt2 := math.Sqrt(float64(number.Number2))
	return c.String(http.StatusOK, "Square root of "+strconv.Itoa(number.Number1)+" is "+strconv.FormatFloat(sqrt1, 'E', -1, 64)+"\nSquare root of "+strconv.Itoa(number.Number2)+" is "+strconv.FormatFloat(sqrt2, 'E', -1, 64))
}
