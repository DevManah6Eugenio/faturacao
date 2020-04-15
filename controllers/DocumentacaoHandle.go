package controllers

import "net/http"

func DocumentacaoHandle(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://app.swaggerhub.com/apis-docs/Manah6EugenioSwagger/Faturacao/1.0.0", http.StatusSeeOther)
}
