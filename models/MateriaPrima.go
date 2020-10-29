package models

type MateriaPrima struct {
	model
	Id            int64        `json:"id" 				 gorm:"Column:id;primaryKey"`
	Codigo        string       `json:"codigo" 			 gorm:"Column:codigo;type:varchar(10)"`
	Nome          string       `json:"nome" 			 gorm:"Column:nome;type:varchar(30)"`
	UnidadeCompra string       `json:"unidadeCompra"	 gorm:"Column:unidade_compra;type:varchar(1)"`
	Valor         float64      `json:"valor"  			 gorm:"Column:valor"`
	Fornecedores  []Fornecedor `json:"fornecedores"`
}

//implementação da inteface do ORM
func (m MateriaPrima) TableName() string {
	return "materia_prima"
}
