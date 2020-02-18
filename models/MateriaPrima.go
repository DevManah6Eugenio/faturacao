package models

type MateriaPrima struct {
	Id            int64        `json:"id"`
	Codigo        string       `json:"codigo"`
	Nome          string       `json:"nome"`
	UnidadeCompra string       `json:"unidadecompra"`
	Valor         float64      `json:"valor"`
	Fornecedores  []Fornecedor `json:"fornecedores"`
}
