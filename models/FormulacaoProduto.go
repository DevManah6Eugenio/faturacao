package models

import "github.com/jinzhu/gorm"

type FormulacaoProduto struct {
	gorm.Model
	Id          int64      `json:"id"  gorm:"Column:id;primary_key"`
	Quantidade  float64    `json:"quantidade"  gorm:"Column:qtd"`
	Valor       float64    `json:"valor"  gorm:"Column:valor"`
	Formulacoes Formulacao `json:"formulacoes" gorm:"foreignkey:Id"`
	produtos    Produto    `json:"produtos" gorm:"foreignkey:Id"`
}

//implementação da inteface do ORM
func (f FormulacaoProduto) TableName() string {
	return "formulacao_produto"
}
