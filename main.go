package main

import (
	"MiniProject/config"
	"miniProject/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
