package models

type FormulacaoProduto struct {
	model
	Quantidade  float64    `json:"quantidade"  	gorm:"Column:qtd"`
	Valor       float64    `json:"valor"  		gorm:"Column:valor"`
	Formulacoes Formulacao `json:"formulacao" 	gorm:"foreignkey:Id"`
	produtos    Produto    `json:"produto" 	    gorm:"foreignkey:Id"`
}

//implementação da inteface do ORM
func (f FormulacaoProduto) TableName() string {
	return "formulacao_produto"
}
