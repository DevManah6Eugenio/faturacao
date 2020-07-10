package models

import "github.com/jinzhu/gorm"

type Formulacao struct {
	gorm.Model
	Id                   int64     `json:"id" gorm:"Column:id;primary_key"`
	Codigo               string    `json:"codigo" gorm:"Column:codigo"`
	QtdEstimadaProduzida float64   `json:"qtdestimadaproduzida"  gorm:"Column:qtd_estimada_produzida"`
	Icm                  float64   `json:"icm"  gorm:"Column:icm"`
	Pis                  float64   `json:"pis"  gorm:"Column:pis"`
	Cofins               float64   `json:"cofins"  gorm:"Column:cofins"`
	Comissao             float64   `json:"comissao"  gorm:"Column:comissao"`
	MargemLucro          float64   `json:"margemlucro"  gorm:"Column:margem_lucro"`
	DespesaFixa          float64   `json:"despesafixa"  gorm:"Column:despesa_fixa"`
	OutrosDescontos      float64   `json:"outrosdescontos"  gorm:"Column:outros_descontos"`
	CustomEstimado       float64   `json:"custoestimado"  gorm:"Column:custo_estimado"`
	Obs                  string    `json:"observacao"  gorm:"Column:obs"`
	ValorVenda           float64   `json:"valorvenda"  gorm:"Column:valor_venda"`
	Produtos             []Produto `json:"produtos"`
}

//implementação da inteface do ORM
func (f Formulacao) TableName() string {
	return "formulacao"
}
