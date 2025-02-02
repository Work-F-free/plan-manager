package main

import "seatPlanner/internal/app"

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
