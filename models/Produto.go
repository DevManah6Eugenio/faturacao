package models

type Produto struct {
	model
	Id            int64          `json:"id" 			gorm:"Column:id"`
	Nome          string         `json:"nome"  			gorm:"Column:nome"`
	Codigo        string         `json:"codigo" 		gorm:"Column:codigo"`
	MateriasPrima []MateriaPrima `json:"materiasPrima" 	gorm:"many2many:produto_materia_prima;"`
}

//implementação da inteface do ORM
func (p Produto) TableName() string {
	return "produto"
}
