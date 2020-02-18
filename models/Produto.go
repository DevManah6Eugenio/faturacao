package models

type Produto struct {
	Id            int64          `json:"id"`
	Nome          string         `json:"nome"`
	Codigo        string         `json:"codigo"`
	MateriasPrima []MateriaPrima `json:"materiasprima"`
}
