package models

type Usuario struct {
	Id    int64  `json:"id"`
	Nome  string `json:"nome"`
	Cpf   string `json:"cpf"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}
