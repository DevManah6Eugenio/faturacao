package dao

import (
	"database/sql"

	"github.com/Manah6Eugenio/faturacao/models"
)

func SelectMateriaPrima(con *sql.DB, filtroMateriaPrima models.MateriaPrima) (materiaPrimas []models.MateriaPrima) {
	selectMateriaPrima := `SELECT id, codigo, nome, unidade_compra, valor 
		FROM materia_prima WHERE (codigo LIKE '%' || $1 || '%') AND (nome LIKE '%' || $2 || '%')`
	rows, err := con.Query(selectMateriaPrima, filtroMateriaPrima.Codigo, filtroMateriaPrima.Nome)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	var m models.MateriaPrima
	for rows.Next() {
		err = rows.Scan(&m.Id, &m.Codigo, &m.Nome, &m.UnidadeCompra, &m.Valor)
		if err != nil {
			panic(err.Error())
		}
		materiaPrimas = append(materiaPrimas, m)
	}
	return
}

func SelectMateriaPrimaId(con *sql.DB, id int64) (materiaPrima models.MateriaPrima) {
	selectMateriaPrimaID := "SELECT id, codigo, nome, unidade_compra, valor FROM materia_prima WHERE id = $1 "
	rows, err := con.Query(selectMateriaPrimaID, id)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	if rows.Next() {
		err = rows.Scan(&materiaPrima.Id, &materiaPrima.Codigo, &materiaPrima.Nome, &materiaPrima.UnidadeCompra, 
			&materiaPrima.Valor)
		if err != nil {
			panic(err.Error())
		}
	}
	return
}

func InsertMateriaPrima(con *sql.DB, materiaPrima *models.MateriaPrima) {
	insertMateriaPrima := "INSERT INTO materia_prima (codigo, nome, unidade_compra, valor) VALUES ($1, $2, $3, $4)"
	tx, _ := con.Begin()
	stmt, _ := tx.Prepare(insertMateriaPrima)
	_, err := stmt.Exec(materiaPrima.Codigo, materiaPrima.Nome, materiaPrima.UnidadeCompra, materiaPrima.Valor)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
}

func UpdateMateriaPrima(con *sql.DB, materiaPrima *models.MateriaPrima) {
	updateMateriaPrima := `UPDATE materia_prima set codigo = $1, nome = $2, unidade_compra = $3, valor = $4 
		WHERE id = $5`
	stmt, _ := con.Prepare(updateMateriaPrima)
	_, err := stmt.Exec(materiaPrima.Codigo, materiaPrima.Nome, materiaPrima.UnidadeCompra, materiaPrima.Valor, 
		materiaPrima.Id)
	if err != nil {
		panic(err)
	}
}

func DeleteMateriaPrima(con *sql.DB, id int64) {
	deleteMateriaPrima := "DELETE FROM materia_prima WHERE id = $1"
	stmt, _ := con.Prepare(deleteMateriaPrima)
	_, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}
