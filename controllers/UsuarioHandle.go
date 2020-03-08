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

func UsuarioHandle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		doGetUsuario(w, r)
	case "POST":
		doPostUsuario(w, r)
	case "PUT":
		doPutUsuario(w, r)
	case "DELETE":
		doDeleteUsuario(w, r)
	case "OPTIONS":
    	utils.ConfigCors(&w);
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func doGetUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var js []byte
	con := db.Connection()
	defer con.Close()
	if r.Form.Get("id") != "" {
		js, _ = json.Marshal(dao.SelectUsuarioId(con, utils.ParseInt(r.Form.Get("id"))))
	} else {
		js, _ = json.Marshal(dao.SelectUsuario(con, 
			models.Usuario{Nome: r.Form.Get("nome"), Cpf: r.Form.Get("cpf")}))
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}

func doPostUsuario(w http.ResponseWriter, r *http.Request) {
	con := db.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var user models.Usuario
	err2 := json.Unmarshal(corpo, &user)
	if err2 != nil {
		panic(err2)
	}
	dao.InsertUsuario(con, &user)
}

func doPutUsuario(w http.ResponseWriter, r *http.Request) {
	con := db.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var user models.Usuario
	err2 := json.Unmarshal(corpo, &user)
	if err2 != nil {
		panic(err2)
	}
	dao.UpdateUsuario(con, &user)
}

func doDeleteUsuario(w http.ResponseWriter, r *http.Request) {
	con := db.Connection()
	defer con.Close()
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var user models.Usuario
	err2 := json.Unmarshal(corpo, &user)
	if err2 != nil {
		panic(err2)
	}
	dao.DeleteUsuario(con, user.Id)
}
