package main

import (
	"log"
	"net/http"

	"github.com/Manah6Eugenio/faturacao/routes"
)

func main() {
	log.Println("servidor iniciado localhost:9000")
	routes.CarregaRotas()
	log.Fatal(http.ListenAndServe(":9000", nil))
}
