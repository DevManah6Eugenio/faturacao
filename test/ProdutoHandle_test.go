package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Manah6Eugenio/faturacao/models"
	"github.com/gorilla/mux"
)

func ProdutoHandle_Test(t *testing.T) {

	r := mux.NewRouter()
	r.Handle("/produto", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := models.Produto{}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
		}
	}))

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
	t.Errorf("deu nem um erro")
}
