package dao

import (
	"log"

	"github.com/Manah6Eugenio/faturacao/models"
	"github.com/jinzhu/gorm"
)

func SelectMateriaPrima(con *gorm.DB, filtroMateriaPrima models.MateriaPrima) (resultQuery []models.MateriaPrima) {
	if filtroMateriaPrima.Id == 0 {
		con.Where(" codigo ILIKE ? OR nome ILIKE ?", "%"+filtroMateriaPrima.Codigo+"%", "%"+filtroMateriaPrima.Nome+"%").Find(&resultQuery)
	} else {
		con.Where(" id = ?", filtroMateriaPrima.Id).Find(&resultQuery)
	}
	return
}

func InsertMateriaPrima(con *gorm.DB, materiaPrima *models.MateriaPrima) models.MateriaPrima {
	if err := con.Create(&materiaPrima).Error; err != nil {
		log.Println(err)
	}
	return *materiaPrima
}

func UpdateMateriaPrima(con *gorm.DB, materiaPrima *models.MateriaPrima) models.MateriaPrima {
	if err := con.Save(&materiaPrima); err != nil {
		log.Println(err)
	}
	return *materiaPrima
}

func DeleteMateriaPrima(con *gorm.DB, id int64) (materiaPrima models.MateriaPrima) {
	con.Where(" id = ?", id).Delete(&materiaPrima)
	return materiaPrima
}
