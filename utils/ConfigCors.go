package utils

import (
	"net/http"
)

func ConfigCors(w *http.ResponseWriter) {
     (*w).Header().Set("Access-Control-Allow-Origin", "*")
     (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
     (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
     (*w).Header().Set("Access-Control-Expose-Headers", "Authorization")
     (*w).Header().Set("Server", "Faturacao")
     (*w).WriteHeader(200)
}