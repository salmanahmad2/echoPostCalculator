package main

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	//_ "github.com/salmanahmad2/echoPostCalculator"

	// _ "github.com/swaggo/echo-swagger/example/docs"
	"github.com/salmanahmad2/echoPostCalculator/controller"
	"github.com/salmanahmad2/echoPostCalculator/database"
	_ "github.com/salmanahmad2/echoPostCalculator/docs"
	"github.com/salmanahmad2/echoPostCalculator/routes"

	_ "github.com/salmanahmad2/echoPostCalculator/routes"
)

var db *sql.DB

func connectToDb() {
	DB := database.DatabaseConnection()

	db = DB
	getDB()
}
func getDB() {
	controller.SetDB(db)
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

	routes.Routes(e)
	connectToDb()
	defer db.Close()
	e.Logger.Fatal(e.Start(":1323"))
	fmt.Println("closed")
}
