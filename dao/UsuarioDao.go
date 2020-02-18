package dao

import (
	"database/sql"

	"github.com/Manah6Eugenio/faturacao/models"
)

func SelectUsuario(con *sql.DB, filtroUser models.Usuario) (users []models.Usuario) {
	selectUser := `SELECT id, nome, cpf, email FROM usuario 
		WHERE (nome LIKE '%' || $1 || '%') AND (cpf LIKE '%' || $2 || '%')`
	rows, err := con.Query(selectUser, filtroUser.Nome, filtroUser.Cpf)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	var user models.Usuario
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Nome, &user.Cpf, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	return
}

func SelectUsuarioId(con *sql.DB, id int64) (user models.Usuario) {
	selectUserID := "SELECT id, nome, cpf, email FROM usuario WHERE id = $1 "
	rows, err := con.Query(selectUserID, id)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Nome, &user.Cpf, &user.Email)
		if err != nil {
			panic(err.Error())
		}
	}
	return
}

func InsertUsuario(con *sql.DB, user *models.Usuario) {
	insertUser := "INSERT INTO usuario (nome, cpf, email, senha) VALUES ($1, $2, $3, $4)"
	tx, _ := con.Begin()
	stmt, _ := tx.Prepare(insertUser)
	_, err := stmt.Exec(user.Nome, user.Cpf, user.Email, user.Senha)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
}

func UpdateUsuario(con *sql.DB, user *models.Usuario) {
	updateUser := "UPDATE usuario set nome = $1, cpf = $2, email = $3, senha = $4 WHERE id = $5"
	stmt, _ := con.Prepare(updateUser)
	_, err := stmt.Exec(user.Nome, user.Cpf, user.Email, user.Senha, user.Id)
	if err != nil {
		panic(err)
	}
}

func DeleteUsuario(con *sql.DB, id int64) {
	deleteUser := "DELETE FROM usuario WHERE id = $1"
	stmt, _ := con.Prepare(deleteUser)
	_, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}
