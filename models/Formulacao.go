package models

type Formulacao struct {
	Id                   int64     `json:"id"`
	Codigo               string    `json:"codigo"`
	QtdEstimadaProduzida float64   `json:"qtdestimadaproduzida"`
	Icm                  float64   `json:"icm"`
	Pis                  float64   `json:"pis"`
	Confis               float64   `json:"confis"`
	Comissao             float64   `json:"comissao"`
	MargemLucro          float64   `json:"margemlucro"`
	DespesaFixa          float64   `json:"despesafixa"`
	OutrosDescontos      float64   `json:"outrosdescontos"`
	CustomEstimado       float64   `json:"custoestimado"`
	Obs                  string    `json:"observacao"`
	ValorVenda           float64   `json:"valorvenda"`
	Produtos             []Produto `json:"produtos"`
}
