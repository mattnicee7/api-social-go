package main

import (
	"api-social-go/src/config"
	"api-social-go/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Println("Rodando API")
	fmt.Println(config.StringConexaoBanco)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
