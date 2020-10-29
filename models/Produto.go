package models

type Produto struct {
	model
	Id            int64          `json:"id" 			gorm:"Column:id;primaryKey"`
	Nome          string         `json:"nome"  			gorm:"Column:nome;type:varchar(30)"`
	Codigo        string         `json:"codigo" 		gorm:"Column:codigo;type:varchar(10)"`
	MateriasPrima []MateriaPrima `json:"materiasPrima" 	gorm:"many2many:produto_materia_prima;"`
}

//implementação da inteface do ORM
func (p Produto) TableName() string {
	return "produto"
}
