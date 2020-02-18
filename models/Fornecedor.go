package models

type Fornecedor struct {
	Id            int64          `json:"id"`
	Cnpj          string         `json:"cnpj"`
	Nome          string         `json:"nome"`
	Empresa       string         `json:"empresa"`
	Telefone1     string         `json:"telefone1"`
	Telefone2     string         `json:"telefone2"`
	Email         string         `json:"email"`
	MateriasPrima []MateriaPrima `json:"materiasprima"`
}
