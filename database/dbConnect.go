package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

func ConnectDataBase() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Falha ao carregar as váriaveis de conexão %v", err)
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
		log.Fatalf("Falha ao tentar conexao com banco de dados %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Falha ao testar comunicação com o banco de dados %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com Sucesso!")
}
