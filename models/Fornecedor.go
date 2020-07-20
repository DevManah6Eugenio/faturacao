package models

type Fornecedor struct {
	model
	Id            int64          `json:"id" 			gorm:"Column:id"`
	Cnpj          string         `json:"cpf_cnpj"  		gorm:"Column:cpf_cnpj"`
	Nome          string         `json:"nome"  			gorm:"Column:nome"`
	Empresa       string         `json:"empresa" 		gorm:"Column:empresa"`
	Telefone      string         `json:"telefone"  	    gorm:"Column:telefone"`
	Celular       string         `json:"celular"  	    gorm:"Column:celular"`
	Email         string         `json:"email"  		gorm:"Column:email"`
	MateriasPrima []MateriaPrima `json:"materiasPrima"  gorm:"many2many:fornecedor_materia_prima;"`
}

//implementação da inteface do ORM
func (f Fornecedor) TableName() string {
	return "fornecedor"
}
