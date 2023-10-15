package main

import (
	"hacktiv8-final-project4/config"
	"hacktiv8-final-project4/route"
)

func main() {
	config.StartDatabase()
	route.StartApplication()
}
