package models

import "github.com/jinzhu/gorm"

type Produto struct {
	gorm.Model
	Id            int64          `json:"id"  gorm:"Column:id;primary_key"`
	Nome          string         `json:"nome"  gorm:"Column:nome"`
	Codigo        string         `json:"codigo"  gorm:"Column:codigo"`
	MateriasPrima []MateriaPrima `json:"materiasprima" gorm:"many2many:produto_materia_prima;"`
}

//implementação da inteface do ORM
func (p Produto) TableName() string {
	return "produto"
}
