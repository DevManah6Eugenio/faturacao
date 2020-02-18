package routes

import (
	"net/http"
	"github.com/Manah6Eugenio/faturacao/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/formulacao", controllers.FormulacaoHandle)
	http.HandleFunc("/fornecedor", controllers.FornecedorHandle)
	http.HandleFunc("/materiaprima", controllers.MateriaPrimaHandle)
	http.HandleFunc("/produto", controllers.ProdutoHandle)
	http.HandleFunc("/usuario", controllers.UsuarioHandle)
}
