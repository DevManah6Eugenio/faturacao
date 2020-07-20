package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Manah6Eugenio/faturacao/dao"
	"github.com/Manah6Eugenio/faturacao/database"
	"github.com/Manah6Eugenio/faturacao/models"
	"github.com/Manah6Eugenio/faturacao/utils"
)

func FormulacaoHandle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		doGetFormulacao(w, r)
	case "POST":
		doPostFormulacao(w, r)
	case "PUT":
		doPutFormulacao(w, r)
	case "DELETE":
		doDeleteFormulacao(w, r)
	case "OPTIONS":
		utils.ConfigCors(&w)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func doGetFormulacao(w http.ResponseWriter, r *http.Request) {
	con := database.Connection()
	defer con.Close()
	r.ParseForm()
	var js []byte
	js, _ = json.Marshal(dao.SelectFormulacao(con,
		models.Formulacao{Obs: r.Form.Get("observacao"), Codigo: r.Form.Get("codigo"), Id: utils.ParseInt(r.Form.Get("id"))}))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}

func doPostFormulacao(w http.ResponseWriter, r *http.Request) {
	con := database.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var f models.Formulacao
	err2 := json.Unmarshal(corpo, &f)
	if err2 != nil {
		panic(err2)
	}
	dao.InsertFormulacao(con, &f)
}

func doPutFormulacao(w http.ResponseWriter, r *http.Request) {
	con := database.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var f models.Formulacao
	err2 := json.Unmarshal(corpo, &f)
	if err2 != nil {
		panic(err2)
	}
	dao.UpdateFormulacao(con, &f)
}

func doDeleteFormulacao(w http.ResponseWriter, r *http.Request) {
	con := database.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var f models.Formulacao
	err2 := json.Unmarshal(corpo, &f)
	if err2 != nil {
		panic(err2)
	}
	dao.DeleteUsuario(con, f.Id)
}
