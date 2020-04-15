package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Manah6Eugenio/faturacao/routes"
)

func main() {
	port := os.Getenv("PORT") //variavel de anbiente com a porta
	if port == "" {
		panic("$PORT não definida")
	}
	log.Println("servidor iniciado na porta: " + port)
	routes.CarregaRotas()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
