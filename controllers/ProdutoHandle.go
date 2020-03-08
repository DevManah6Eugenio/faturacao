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

func ProdutoHandle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		doGetProduto(w, r)
	case "POST":
		doPostProduto(w, r)
	case "PUT":
		doPutProduto(w, r)
	case "DELETE":
		doDeleteProduto(w, r)
	case "OPTIONS":
    	utils.ConfigCors(&w)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func doGetProduto(w http.ResponseWriter, r *http.Request) {
	con := db.Connection()
	defer con.Close()
	r.ParseForm()
	var js []byte
	if r.Form.Get("id") != "" {
		js, _ = json.Marshal(dao.SelectProdutoId(con, utils.ParseInt(r.Form.Get("id"))))
	} else {
		js, _ = json.Marshal(dao.SelectProduto(con, 
			models.Produto{Nome: r.Form.Get("nome"), Codigo: r.Form.Get("codigo")}))
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}

func doPostProduto(w http.ResponseWriter, r *http.Request) {
	con := db.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var p models.Produto
	err2 := json.Unmarshal(corpo, &p)
	if err2 != nil {
		panic(err2)
	}
	dao.InsertProduto(con, &p)
}

func doPutProduto(w http.ResponseWriter, r *http.Request) {
	con := db.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var p models.Produto
	err2 := json.Unmarshal(corpo, &p)
	if err2 != nil {
		panic(err2)
	}
	dao.UpdateProduto(con, &p)
}

func doDeleteProduto(w http.ResponseWriter, r *http.Request) {
	con := db.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var p models.Produto
	err2 := json.Unmarshal(corpo, &p)
	if err2 != nil {
		panic(err2)
	}
	dao.DeleteProduto(con, p.Id)
}
