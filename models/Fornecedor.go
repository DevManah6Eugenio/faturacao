package models

type Fornecedor struct {
	model
	Id            int64          `json:"id" 			gorm:"Column:id;primaryKey"`
	Cnpj          string         `json:"cpf_cnpj"  		gorm:"Column:cpf_cnpj;type:varchar(14)"`
	Nome          string         `json:"nome"  			gorm:"Column:nome;type:varchar(30)"`
	Empresa       string         `json:"empresa" 		gorm:"Column:empresa;type:varchar(30)"`
	Telefone      string         `json:"telefone"  	    gorm:"Column:telefone;type:varchar(16)"`
	Celular       string         `json:"celular"  	    gorm:"Column:celular;type:varchar(16)"`
	Email         string         `json:"email"  		gorm:"Column:email;type:varchar(30)"`
	MateriasPrima []MateriaPrima `json:"materiasPrima"  gorm:"many2many:fornecedor_materia_prima;"`
}

//implementação da inteface do ORM
func (f Fornecedor) TableName() string {
	return "fornecedor"
}
