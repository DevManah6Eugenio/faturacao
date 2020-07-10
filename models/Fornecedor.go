package models

import "github.com/jinzhu/gorm"

type Fornecedor struct {
	gorm.Model
	Id            int64          `json:"id"  gorm:"Column:id;primary_key"`
	Cnpj          string         `json:"cnpj"  gorm:"Column:cpf_cnpj"`
	Nome          string         `json:"nome"  gorm:"Column:nome"`
	Empresa       string         `json:"empresa"  gorm:"Column:empresa"`
	Telefone1     string         `json:"telefone1"  gorm:"Column:tel01"`
	Telefone2     string         `json:"telefone2"  gorm:"Column:tel02"`
	Email         string         `json:"email"  gorm:"Column:email"`
	MateriasPrima []MateriaPrima `json:"materiasprima" gorm:"many2many:fornecedor_materia_prima;"`
}

//implementação da inteface do ORM
func (f Fornecedor) TableName() string {
	return "fornecedor"
}
