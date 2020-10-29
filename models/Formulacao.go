package models

type Formulacao struct {
	model
	Id                   int64               `json:"id" 					    gorm:"Column:id;primaryKey"`
	Codigo               string              `json:"codigo" 					gorm:"Column:codigo;type:varchar(10)"`
	QtdEstimadaProduzida float64             `json:"qtdEstimadaProduzida"    	gorm:"Column:qtd_estimada_produzida"`
	Icm                  float64             `json:"icm"  				    	gorm:"Column:icm"`
	Pis                  float64             `json:"pis"  					    gorm:"Column:pis"`
	Cofins               float64             `json:"cofins"  					gorm:"Column:cofins"`
	Comissao             float64             `json:"comissao"  			    	gorm:"Column:comissao"`
	MargemLucro          float64             `json:"margemLucro"  		    	gorm:"Column:margem_lucro"`
	DespesaFixa          float64             `json:"despesaFixa"  		    	gorm:"Column:despesa_fixa"`
	OutrosDescontos      float64             `json:"outrosDescontos"  	    	gorm:"Column:outros_descontos"`
	CustomEstimado       float64             `json:"custoEstimado"  			gorm:"Column:custo_estimado"`
	Obs                  string              `json:"observacao"  				gorm:"Column:obs;type:varchar(50)"`
	ValorVenda           float64             `json:"valorVenda"  				gorm:"Column:valor_venda"`
	FormulacaoProduto    []FormulacaoProduto `json:"formulacaoProduto"`
}

//implementação da inteface do ORM
func (f Formulacao) TableName() string {
	return "formulacao"
}
