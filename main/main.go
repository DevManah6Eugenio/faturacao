package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Manah6Eugenio/faturacao/routes"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT não definida")
	}

	log.Println("servidor iniciado localhost:" + port)
	routes.CarregaRotas()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
