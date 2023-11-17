package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	utils.CarregarTemplates()
	rotas := router.Gerar()

	fmt.Println(config.HashKey)

	fmt.Printf("Escutando na porta: %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), rotas))
}
