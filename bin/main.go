package main

import (
	"fmt"
	"log"
	"os"

	"treinamento-api/routes"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func connectDataBase() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("não pode carregar as váriaveis de conexão %v", err)
	}

	dbUser := os.Getenv("USUARIO")
	dbPass := os.Getenv("SENHA")
	dbHost := os.Getenv("SERVIDOR")
	dbPort := os.Getenv("PORTA")
	dbName := os.Getenv("NOMEDB")

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	var err error

	DB, err = sqlx.Connect("postgres", dataSource)

	if err != nil {
		log.Fatalf("não pode conectar no banco de dados %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("falha ao testar comunicação com o banco de dados %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com Sucesso!")

}

func main() {

	fmt.Println("Hello World")
	connectDataBase()
	routes.GetRoutes()
}
