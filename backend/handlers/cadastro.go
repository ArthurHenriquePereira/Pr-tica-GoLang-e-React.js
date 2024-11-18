package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend/db" // Caminho correto para o pacote db
)

// Função para salvar os dados no banco de dados
func salvarCadastro(dados map[string]interface{}) error {
	// Converta os dados do map para variáveis específicas
	email, okEmail := dados["email"].(string)
	senha, okSenha := dados["senha"].(string)

	if !okEmail || !okSenha {
		return fmt.Errorf("dados inválidos")
	}

	// Query de inserção no banco de dados
	query := "INSERT INTO cadastros (email, senha) VALUES (?, ?)"
	_, err := db.DB.Exec(query, email, senha)
	return err
}

// Handler para o cadastro
func CadastroHandler(w http.ResponseWriter, r *http.Request) {
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

		// Salvar os dados no banco de dados
		if err := salvarCadastro(dados); err != nil {
			http.Error(w, "Erro ao salvar os dados no banco de dados", http.StatusInternalServerError)
			return
		}

		// Responder com uma mensagem de sucesso
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "Cadastro recebido com sucesso"}`)
	} else {
		// Se a requisição não for POST
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
