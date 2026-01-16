package main

import (
	"api-social-go/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("oi")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
