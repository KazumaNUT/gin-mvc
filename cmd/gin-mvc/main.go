package main

import (
	"github.com/gin-mvc/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}
