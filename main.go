package main

import (
	usersController "fluke-hospital/controller"
	"fluke-hospital/db"
	"fmt"
)

var client = db.DbConnect()

func main() {
	fmt.Println("Fluke Hospital go is running!!!")
	fmt.Println("list user:", usersController.ListUser())
}
