package main

import (
	"rest-api/app/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}
