package dao

import (
	"github.com/Manah6Eugenio/faturacao/models"
	"github.com/jinzhu/gorm"
)

func SelectProduto(con *gorm.DB, filtroProduto models.Produto) (produtos []models.Produto) {
	// selectProduto := `SELECT id, nome, codigo FROM produto
	// 	WHERE (nome LIKE '%' || $1 || '%') AND (codigo LIKE '%' || $2 || '%') `
	// rows, err := con.Query(selectProduto, filtroProduto.Nome, filtroProduto.Codigo)
	// defer rows.Close()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// var p models.Produto
	// for rows.Next() {
	// 	err = rows.Scan(&p.Id, &p.Nome, &p.Codigo)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	produtos = append(produtos, p)
	// }
	return
}

func SelectProdutoId(con *gorm.DB, id int64) (produto models.Produto) {
	// selectProdutoID := "SELECT id, nome, codigo FROM produto WHERE id = $1 "
	// rows, err := con.Query(selectProdutoID, id)
	// defer rows.Close()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// if rows.Next() {
	// 	err = rows.Scan(&produto.Id, &produto.Nome, &produto.Codigo)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }
	return
}

func InsertProduto(con *gorm.DB, produto *models.Produto) {
	// insertProduto := "INSERT INTO produto (nome, codigo) VALUES ($1, $2) "
	// tx, _ := con.Begin()
	// stmt, _ := tx.Prepare(insertProduto)
	// _, err := stmt.Exec(produto.Nome, produto.Codigo)
	// if err != nil {
	// 	tx.Rollback()
	// 	panic(err)
	// }
	// tx.Commit()
}

func UpdateProduto(con *gorm.DB, produto *models.Produto) {
	// updateProduto := "UPDATE produto set nome = $1, codigo = $2 WHERE id = $3 "
	// stmt, _ := con.Prepare(updateProduto)
	// _, err := stmt.Exec(produto.Nome, produto.Codigo, produto.Id)
	// if err != nil {
	// 	panic(err)
	// }
}

func DeleteProduto(con *gorm.DB, id int64) {
	// deleteProduto := "DELETE FROM produto WHERE id = $1 "
	// stmt, _ := con.Prepare(deleteProduto)
	// _, err := stmt.Exec(id)
	// if err != nil {
	// 	panic(err)
	// }
}
