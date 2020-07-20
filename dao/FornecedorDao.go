package dao

import (
	"log"

	"github.com/Manah6Eugenio/faturacao/models"
	"github.com/jinzhu/gorm"
)

func SelectFornecedor(con *gorm.DB, filtroFornecedor models.Fornecedor) (resultQuery []models.Fornecedor) {
	if filtroFornecedor.Id == 0 {
		con.Where(" nome ILIKE ? OR cpf_cnpj ILIKE ?", "%"+filtroFornecedor.Nome+"%", "%"+filtroFornecedor.Cnpj+"%").Find(&resultQuery)
	} else {
		con.Where(" id = ?", filtroFornecedor.Id).Find(&resultQuery)
	}
	return
}

func InsertFornecedor(con *gorm.DB, fornecedor *models.Fornecedor) models.Fornecedor {
	if err := con.Create(&fornecedor).Error; err != nil {
		log.Println(err)
	}
	return *fornecedor
}

func UpdateFornecedor(con *gorm.DB, fornecedor *models.Fornecedor) models.Fornecedor {
	if err := con.Save(&fornecedor); err != nil {
		log.Println(err)
	}
	return *fornecedor
}

func DeleteFornecedor(con *gorm.DB, id int64) (fornecedor models.Fornecedor) {
	con.Where(" id = ?", id).Delete(&fornecedor)
	return fornecedor
}
