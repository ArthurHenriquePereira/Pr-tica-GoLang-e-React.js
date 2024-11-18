package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Variável global para manter a conexão com o banco de dados
var DB *sql.DB

// Função para criar e configurar a conexão com o banco de dados
func InitDB() error {
	// Substitua com suas credenciais de banco de dados
	dsn := "root:@tcp(localhost:3307)/compras?parseTime=true"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	// Testar a conexão
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("erro ao testar conexão com o banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso.")
	return nil
}

// Função para fechar a conexão com o banco de dados
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
