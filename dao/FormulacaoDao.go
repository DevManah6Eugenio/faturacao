package dao

import (
	"log"

	"github.com/Manah6Eugenio/faturacao/models"
	"github.com/jinzhu/gorm"
)

func SelectFormulacao(con *gorm.DB, filtroFormulacao models.Formulacao) (resultQuery []models.Formulacao) {
	if filtroFormulacao.Id == 0 {
		con.Where(" observacao ILIKE ? OR codigo ILIKE ?", "%"+filtroFormulacao.Obs+"%", "%"+filtroFormulacao.Codigo+"%").Find(&resultQuery)
	} else {
		con.Where(" id = ?", filtroFormulacao.Id).Find(&resultQuery)
	}
	return
}

func InsertFormulacao(con *gorm.DB, formulacao *models.Formulacao) models.Formulacao {
	if err := con.Create(&formulacao).Error; err != nil {
		log.Println(err)
	}
	return *formulacao
}

func UpdateFormulacao(con *gorm.DB, formulacao *models.Formulacao) models.Formulacao {
	if err := con.Save(&formulacao); err != nil {
		log.Println(err)
	}
	return *formulacao
}

func DeleteFormulacao(con *gorm.DB, id int64) (formulacao models.Formulacao) {
	con.Where(" id = ?", id).Delete(&formulacao)
	return formulacao
}
