package main

import (
	"Project-Mini/config"
	"Project-Mini/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8090")
}
