package routes

import (
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Numbers struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type Response struct {
	Result float64 `json:"result"`
}

func Add(c echo.Context) error {
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

func Sub(c echo.Context) error {
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

func Multiply(c echo.Context) error {
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

func Division(c echo.Context) error {
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

func Modulus(c echo.Context) error {
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

func Square(c echo.Context) error {
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

func Power(c echo.Context) error {
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

func Sqrt(c echo.Context) error {
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
