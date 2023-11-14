package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTemplates()
	rotas := router.Gerar()

	fmt.Println("Rodando o web-app")
	log.Fatal(http.ListenAndServe(":3000", rotas))
}
