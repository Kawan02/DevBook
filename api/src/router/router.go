package router

import (
	"api/src/router/rotas"
	"fmt"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	fmt.Println("Rodando a API!")
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
