package main

import (
	"log"

	"treinamento-api/database"
	"treinamento-api/routes"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Conectando no banco de dados...")
	database.ConnectDataBase()
	routes.GetRoutes()
}
