package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Manah6Eugenio/faturacao/controllers"
)

func main() {
	port := os.Getenv("PORT") //variavel de anbiente com a porta
	if port == "" {
		panic("$PORT não definida")
	}
	log.Println("servidor iniciado na porta: " + port)
	carregaRotas()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func carregaRotas() {
	http.HandleFunc("/formulacao", controllers.FormulacaoHandle)
	http.HandleFunc("/fornecedor", controllers.FornecedorHandle)
	http.HandleFunc("/materiaprima", controllers.MateriaPrimaHandle)
	http.HandleFunc("/produto", controllers.ProdutoHandle)
	http.HandleFunc("/usuario", controllers.UsuarioHandle)
	http.HandleFunc("/", documentacaoHandle) //documentação
}

func documentacaoHandle(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://app.swaggerhub.com/apis-docs/Manah6EugenioSwagger/Faturacao/1.0.0", http.StatusSeeOther)
}
