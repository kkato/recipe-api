package main

import (
	"github.com/kkato/recipe-api/app/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}
