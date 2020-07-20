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

func MateriaPrimaHandle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		doGetMateriaPrima(w, r)
	case "POST":
		doPostMateriaPrima(w, r)
	case "PUT":
		doPutMateriaPrima(w, r)
	case "DELETE":
		doDeleteMateriaPrima(w, r)
	case "OPTIONS":
		utils.ConfigCors(&w)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func doGetMateriaPrima(w http.ResponseWriter, r *http.Request) {
	con := database.Connection()
	defer con.Close()
	r.ParseForm()
	var js []byte
	js, _ = json.Marshal(dao.SelectMateriaPrima(con,
		models.MateriaPrima{Codigo: r.Form.Get("codigo"), Nome: r.Form.Get("nome"), Id: utils.ParseInt(r.Form.Get("id"))}))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}

func doPostMateriaPrima(w http.ResponseWriter, r *http.Request) {
	con := database.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var m models.MateriaPrima
	err2 := json.Unmarshal(corpo, &m)
	if err2 != nil {
		panic(err2)
	}
	dao.InsertMateriaPrima(con, &m)
}

func doPutMateriaPrima(w http.ResponseWriter, r *http.Request) {
	con := database.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var m models.MateriaPrima
	err2 := json.Unmarshal(corpo, &m)
	if err2 != nil {
		panic(err2)
	}
	dao.UpdateMateriaPrima(con, &m)
}

func doDeleteMateriaPrima(w http.ResponseWriter, r *http.Request) {
	con := database.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var m models.MateriaPrima
	err2 := json.Unmarshal(corpo, &m)
	if err2 != nil {
		panic(err2)
	}
	dao.DeleteMateriaPrima(con, m.Id)
}
