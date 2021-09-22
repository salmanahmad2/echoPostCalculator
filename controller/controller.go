package controller

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	//"github.com/salmanahmad2/echoPostCalculator/database"
	_ "github.com/salmanahmad2/echoPostCalculator/database"

	"github.com/labstack/echo/v4"
)

var check int
var db *sql.DB

func SetDB(DB *sql.DB) {
	db = DB
}

type Numbers struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type Response struct {
	Result float64 `json:"result"`
}

// type getId struct {
// 	ID int `json:"id"`
// }
type fetchedData struct {
	Id      int     `json:"id"`
	Num1    float64 `json:"num1"`
	Num2    float64 `json:"num2"`
	Opr     string  `json:"operation"`
	Rslt    float64 `json:"result"`
	Created string  `json:"created at"`
}

func (number Numbers) insertIntoDatabase(result float64, operation string) {

	//defer db.Close()
	sql := "INSERT INTO calculate(number1, number2, operation, result) VALUES( ?, ?,?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(number.Number1, number.Number2, operation, result)

	if err2 != nil {
		panic(err2)
	}
}

func GetRecord(c echo.Context) error {
	fmt.Println("HEADER")
	fmt.Println(c.Request().Header["Authorization"])
	var number1 float64
	var ID int
	var number2 float64
	var Operation string
	var Result float64
	var createdAt string

	id := c.Param("id")
	fmt.Println(id)

	Err := db.QueryRow("SELECT * FROM calculate WHERE ID = ?", id).Scan(&ID, &number1, &number2, &Operation, &Result, &createdAt)
	if Err != nil {
		fmt.Println(Err.Error())
	}
	response := fetchedData{Id: ID, Num1: number1, Num2: number2, Opr: Operation, Rslt: Result, Created: createdAt}
	fmt.Println(response)
	return c.JSON(http.StatusOK, response)
}

func GetAllRecords(c echo.Context) error {
	var number1 float64
	var ID int
	var number2 float64
	var Operation string
	var Result float64
	var createdAt string

	respData := make([]fetchedData, 0)
	rows, err := db.Query("SELECT * FROM calculate")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&ID, &number1, &number2, &Operation, &Result, &createdAt); err != nil {
			panic(err)
		}
		record := fetchedData{Id: ID, Num1: number1, Num2: number2, Opr: Operation, Rslt: Result, Created: createdAt}

		respData = append(respData, record)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, respData)
}

func DeleteRecord(c echo.Context) error {
	id := c.Param("id")
	_, err := db.Query("DELETE FROM calculate where id=?", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c.JSON(http.StatusOK, "Success")
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
	number.insertIntoDatabase(add, "+")

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
	number.insertIntoDatabase(sub, "-")

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
	number.insertIntoDatabase(multiply, "*")
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
	number.insertIntoDatabase(division, "/")
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
	number.insertIntoDatabase(float64(modulus), "%")
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
	number.insertIntoDatabase(square1, "square")
	return c.JSON(http.StatusOK, result)
}

func Power(c echo.Context) error {
	number := new(Numbers)
	if err := c.Bind(number); err != nil {
		return err
	}
	fmt.Println(number)
	power := math.Pow(number.Number1, number.Number2)
	result := Response{
		power,
	}
	number.insertIntoDatabase(power, "power")
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
	number.insertIntoDatabase(sqrt1, "square root")
	return c.JSON(http.StatusOK, result)
	//hello

}
