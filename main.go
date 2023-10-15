package main

import (
	"hacktiv8-final-project4/config"
	"hacktiv8-final-project4/routes"
)

func main() {
	config.StartDatabase()
	routes.StartApplication()
}
