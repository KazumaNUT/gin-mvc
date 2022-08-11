package main

import (
	"gin-mvc/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}
