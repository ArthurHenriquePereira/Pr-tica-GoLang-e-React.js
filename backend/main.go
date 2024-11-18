package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func cadastroHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar se é uma requisição POST
	if r.Method == http.MethodPost {
		// Acessando o corpo da requisição (para garantir que os dados estão chegando corretamente)
		var dados map[string]interface{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&dados); err != nil {
			http.Error(w, "Erro ao ler os dados", http.StatusBadRequest)
			return
		}

		// Imprimir os dados recebidos para debug
		fmt.Printf("Dados recebidos: %+v\n", dados)

		// Responder com uma mensagem de sucesso
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "Cadastro recebido com sucesso"}`)
	} else {
		// Se a requisição não for POST
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Configuração do CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Permite o frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Roteamento
	http.HandleFunc("/api/cadastro", cadastroHandler)

	// Aplica o CORS e inicia o servidor
	handler := c.Handler(http.DefaultServeMux)
	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
