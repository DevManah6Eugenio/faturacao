package models

import "github.com/jinzhu/gorm"

type Usuario struct {
	gorm.Model
	Id    int64  `json:"id"  gorm:"Column:id;primary_key"`
	Nome  string `json:"nome"  gorm:"Column:nome"`
	Cpf   string `json:"cpf"  gorm:"Column:cpf_cnpj"`
	Email string `json:"email"  gorm:"Column:email"`
	Senha string `json:"senha"  gorm:"Column:senha"`
}

//implementação da inteface do ORM
func (p Usuario) TableName() string {
	return "usuario"
}
