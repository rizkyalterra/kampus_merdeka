package main

import (
	"kampus_merdeka/configs"
	"kampus_merdeka/routes"
)

func main() {
	configs.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
}
