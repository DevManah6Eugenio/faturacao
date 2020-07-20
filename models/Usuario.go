package models

type Usuario struct {
	model
	Id    int64  `json:"id" 		gorm:"Column:id"`
	Nome  string `json:"nome"   	gorm:"Column:nome;type:varchar(30);not null;"`
	Cpf   string `json:"cpf_cnpj"   gorm:"Column:cpf_cnpj;type:varchar(14);not null;"`
	Email string `json:"email"  	gorm:"Column:email;type:varchar(30);not null;"`
	Senha string `json:"senha" 		gorm:"Column:senha;type:varchar(30);not null;"`
}

//implementação da inteface do ORM
func (p Usuario) TableName() string {
	return "usuario"
}
