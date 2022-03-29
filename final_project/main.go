package main

import (
	"github.com/MuhGhifari/GolangBootcamp/final-project/config"
	"github.com/MuhGhifari/GolangBootcamp/final-project/router"
)

func main() {

	config.StartDB()

	r := router.StartApp()

	r.Run(":8080")
}
