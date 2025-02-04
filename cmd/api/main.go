package main

import (
	_ "seatPlanner/docs"
	"seatPlanner/internal/app"
)

// @title           Coworking Plan Storage
// @version         1.0
// @description     Service for storing Coworking Plan Data

// @contact.name   Koreshkov Daniil
// @contact.email  danielkoreshkov@gmail.com

// @host      localhost:8080
// @BasePath /
func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
