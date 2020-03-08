package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Manah6Eugenio/faturacao/dao"
	"github.com/Manah6Eugenio/faturacao/db"
	"github.com/Manah6Eugenio/faturacao/models"
	"github.com/Manah6Eugenio/faturacao/utils"
)

func FornecedorHandle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		doGetFornecedor(w, r)
	case "POST":
		doPostFornecedor(w, r)
	case "PUT":
		doPutFornecedor(w, r)
	case "DELETE":
		doDeleteFornecedor(w, r)
	case "OPTIONS":
    	utils.ConfigCors(&w);
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func doGetFornecedor(w http.ResponseWriter, r *http.Request) {
	con := db.Connection()
	defer con.Close()
	r.ParseForm()
	var js []byte
	if r.Form.Get("id") != "" {
		js, _ = json.Marshal(dao.SelectFornecedorId(con, utils.ParseInt(r.Form.Get("id"))))
	} else {
		js, _ = json.Marshal(dao.SelectFornecedor(con, 
			models.Fornecedor{Nome: r.Form.Get("nome"), Cnpj: r.Form.Get("cnpj")}))
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}

func doPostFornecedor(w http.ResponseWriter, r *http.Request) {
	con := db.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var f models.Fornecedor
	err2 := json.Unmarshal(corpo, &f)
	if err2 != nil {
		panic(err2)
	}
	dao.InsertFornecedor(con, &f)
}

func doPutFornecedor(w http.ResponseWriter, r *http.Request) {
	con := db.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var f models.Fornecedor
	err2 := json.Unmarshal(corpo, &f)
	if err2 != nil {
		panic(err2)
	}
	dao.UpdateFornecedor(con, &f)
}

func doDeleteFornecedor(w http.ResponseWriter, r *http.Request) {
	con := db.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var f models.Fornecedor
	err2 := json.Unmarshal(corpo, &f)
	if err2 != nil {
		panic(err2)
	}
	dao.DeleteFornecedor(con, f.Id)
}
