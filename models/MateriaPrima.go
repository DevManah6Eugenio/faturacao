package models

type MateriaPrima struct {
	model
	Id            int64        `json:"id" 				 gorm:"Column:id"`
	Codigo        string       `json:"codigo" 			 gorm:"Column:codigo"`
	Nome          string       `json:"nome" 			 gorm:"Column:nome"`
	UnidadeCompra string       `json:"unidadeCompra"	 gorm:"Column:unidade_compra"`
	Valor         float64      `json:"valor"  			 gorm:"Column:valor"`
	Fornecedores  []Fornecedor `json:"fornecedores"`
}

//implementação da inteface do ORM
func (m MateriaPrima) TableName() string {
	return "materia_prima"
}
