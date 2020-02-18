package models

type FormulacaoProduto struct {
	Id          int64        `json:"id"`
	Quantidade  float64      `json:"quantidade"`
	Valor       float64      `json:"valor"`
	Formulacoes []Formulacao `json:"formulacoes"`
	produtos    []Produto    `json:"produtos"`
}
