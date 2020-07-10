package dao

import (
	"github.com/Manah6Eugenio/faturacao/models"
	"github.com/jinzhu/gorm"
)

func SelectFornecedor(con *gorm.DB, filtroFornecedor models.Fornecedor) (fornecedores []models.Fornecedor) {
	// selectFornecedor := `SELECT  id, cnpj, nome, empresa, telefone1, telefone2, email
	// 	FROM fornecedor WHERE (cnpj ilike '%' || $1 || '%') AND (nome ilike '%' || $2 || '%')`
	// rows, err := con.Query(selectFornecedor, filtroFornecedor.Cnpj, filtroFornecedor.Nome)
	// defer rows.Close()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// var f models.Fornecedor
	// for rows.Next() {
	// 	err = rows.Scan(&f.Id, &f.Cnpj, &f.Nome, &f.Empresa, &f.Telefone1, &f.Telefone2, &f.Email)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fornecedores = append(fornecedores, f)
	// }
	return
}

func SelectFornecedorId(con *gorm.DB, id int64) (fornecedor models.Fornecedor) {
	// selectFornecedorId := `SELECT id, cnpj, nome, empresa, telefone1, telefone2, email
	// 	FROM fornecedor WHERE id = $1 `
	// rows, err := con.Query(selectFornecedorId, id)
	// defer rows.Close()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// if rows.Next() {
	// 	err = rows.Scan(&fornecedor.Id, &fornecedor.Cnpj, &fornecedor.Nome, &fornecedor.Empresa,
	// 		&fornecedor.Telefone1, &fornecedor.Telefone2, &fornecedor.Email)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }
	return
}

func InsertFornecedor(con *gorm.DB, fornecedor *models.Fornecedor) {
	// insertFornecedor := `INSERT INTO fornecedor (cnpj, nome, empresa, telefone1, telefone2, email)
	// 	VALUES ( $1, $2, $3, $4, $5, $6)`
	// tx, _ := con.Begin()
	// stmt, _ := tx.Prepare(insertFornecedor)
	// _, err := stmt.Exec(fornecedor.Cnpj, fornecedor.Nome, fornecedor.Empresa, fornecedor.Telefone1,
	// 	fornecedor.Telefone2, fornecedor.Email)
	// if err != nil {
	// 	tx.Rollback()
	// 	panic(err)
	// }
	// tx.Commit()
}

func UpdateFornecedor(con *gorm.DB, fornecedor *models.Fornecedor) {
	// updateFornecedor := `UPDATE fornecedor set cnpj = $1, nome = $2, empresa = $3, telefone1 = $4,
	// 	telefone2 = $5, email = $6 WHERE id = $7 `
	// stmt, _ := con.Prepare(updateFornecedor)
	// _, err := stmt.Exec(fornecedor.Cnpj, fornecedor.Nome, fornecedor.Empresa, fornecedor.Telefone1,
	// 	fornecedor.Telefone2, fornecedor.Email, fornecedor.Id)
	// if err != nil {
	// 	panic(err)
	// }
}

func DeleteFornecedor(con *gorm.DB, id int64) {
	// deleteFornecedor := "DELETE FROM fornecedor WHERE id = $1"
	// stmt, _ := con.Prepare(deleteFornecedor)
	// _, err := stmt.Exec(id)
	// if err != nil {
	// 	panic(err)
	// }
}
