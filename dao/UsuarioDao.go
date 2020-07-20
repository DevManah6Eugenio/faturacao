package dao

import (
	"log"

	"github.com/Manah6Eugenio/faturacao/models"
	"github.com/jinzhu/gorm"
)

func SelectUsuario(con *gorm.DB, filtroUser models.Usuario) (resultQuery []models.Usuario) {
	if filtroUser.Id == 0 {
		con.Where(" cpf_cnpj ILIKE ? OR nome ILIKE ?", "%"+filtroUser.Cpf+"%", "%"+filtroUser.Nome+"%").Find(&resultQuery)
	} else {
		con.Where(" id = ?", filtroUser.Id).Find(&resultQuery)
	}
	return
}

func InsertUsuario(con *gorm.DB, user *models.Usuario) models.Usuario {
	if err := con.Create(&user).Error; err != nil {
		log.Println(err)
	}
	return *user
}

func UpdateUsuario(con *gorm.DB, user *models.Usuario) models.Usuario {
	if err := con.Save(&user); err != nil {
		log.Println(err)
	}
	return *user
}

func DeleteUsuario(con *gorm.DB, id int64) (user models.Usuario) {
	con.Where(" id = ?", id).Delete(&user)
	return user
}
