package main

import (
	"backend/db"
	"backend/handlers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Inicializar o banco de dados
	if err := db.InitDB(); err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.CloseDB() // Garantir que a conexão será fechada

	// Configuração do CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Permite o frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Roteamento
	http.HandleFunc("/api/cadastro", handlers.CadastroHandler)

	// Aplica o CORS e inicia o servidor
	handler := c.Handler(http.DefaultServeMux)
	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
