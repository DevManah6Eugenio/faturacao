package dao

import (
	"database/sql"

	"github.com/Manah6Eugenio/faturacao/models"
)

func SelectFormulacao(con *sql.DB, filtroFormulacao models.Formulacao) (formulacoes []models.Formulacao) {
	selectFormulacao := `SELECT id, qtd_estimada_produzida, icm, pis, cofins, comissao, 
		margem_lucro, despesa_fixa, outros_descontos, custo_estimado, valor_venda, codigo, obs
		FROM formulacao WHERE (codigo ILIKE '%' || $1 || '%') or (obs ILIKE '%' || $2 || '%') `
	rows, err := con.Query(selectFormulacao, filtroFormulacao.Codigo, filtroFormulacao.Obs)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	var f models.Formulacao
	for rows.Next() {
		err = rows.Scan(&f.Id, &f.QtdEstimadaProduzida, &f.Icm, &f.Pis, &f.Confis, &f.Comissao, 
			&f.MargemLucro, &f.DespesaFixa, &f.OutrosDescontos, &f.CustomEstimado, &f.ValorVenda, 
			&f.Codigo, &f.Obs)
		if err != nil {
			panic(err.Error())
		}
		formulacoes = append(formulacoes, f)
	}
	return
}

func SelectFormulacaoId(con *sql.DB, id int64) (formulacao models.Formulacao) {
	selectFormulacaoID := `SELECT id, qtd_estimada_produzida, icm, pis, cofins, comissao, 
		margem_lucro, despesa_fixa, outros_descontos, custo_estimado, valor_venda, codigo, obs
		FROM formulacao WHERE id = $1 `
	rows, err := con.Query(selectFormulacaoID, id)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	if rows.Next() {
		err = rows.Scan(&formulacao.Id, &formulacao.QtdEstimadaProduzida, &formulacao.Icm, &formulacao.Pis, 
			&formulacao.Confis, &formulacao.Comissao, &formulacao.MargemLucro, &formulacao.DespesaFixa, 
			&formulacao.OutrosDescontos, &formulacao.CustomEstimado, &formulacao.ValorVenda, 
			&formulacao.Codigo, &formulacao.Obs)
		if err != nil {
			panic(err.Error())
		}
	}
	return
}

func InsertFormulacao(con *sql.DB, formulacao *models.Formulacao) {
	insertFormulacao := `INSERT INTO formulacao (qtd_estimada_produzida, icm, pis, cofins, 
		comissao, margem_lucro, despesa_fixa, outros_descontos, custo_estimado, valor_venda, codigo, obs) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	tx, _ := con.Begin()
	stmt, _ := tx.Prepare(insertFormulacao)
	_, err := stmt.Exec(&formulacao.QtdEstimadaProduzida, &formulacao.Icm, &formulacao.Pis, &formulacao.Confis, 
		&formulacao.Comissao, &formulacao.MargemLucro, &formulacao.DespesaFixa, &formulacao.OutrosDescontos, 
		&formulacao.CustomEstimado, &formulacao.ValorVenda, &formulacao.Codigo, &formulacao.Obs)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
}

func UpdateFormulacao(con *sql.DB, formulacao *models.Formulacao) {
	updateFormulacao := `UPDATE formulacao set qtd_estimada_produzida = $1, icm = $2, pis = $3, cofins = $4, 
		comissao = $5, margem_lucro = $6, despesa_fixa = $7, outros_descontos = $8, custo_estimado = $9, 
		valor_venda = $10, codigo = $11, obs = $12 WHERE id = $13`
	stmt, _ := con.Prepare(updateFormulacao)
	_, err := stmt.Exec(formulacao.QtdEstimadaProduzida, formulacao.Icm, formulacao.Pis, formulacao.Confis, 
		formulacao.Comissao, formulacao.MargemLucro, formulacao.DespesaFixa, formulacao.OutrosDescontos, 
		formulacao.CustomEstimado, formulacao.ValorVenda, formulacao.Codigo, formulacao.Obs, formulacao.Id)
	if err != nil {
		panic(err)
	}
}

func DeleteFormulacao(con *sql.DB, id int64) {
	deleteFormulacao := "DELETE FROM formulacao WHERE id = $1"
	stmt, _ := con.Prepare(deleteFormulacao)
	_, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}
