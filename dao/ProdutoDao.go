package dao

import (
	"log"

	"github.com/Manah6Eugenio/faturacao/models"
	"github.com/jinzhu/gorm"
)

func SelectProduto(con *gorm.DB, filtroProduto models.Produto) (resultQuery []models.Produto) {
	if filtroProduto.Id == 0 {
		con.Where(" nome ILIKE ? OR codigo ILIKE ?", "%"+filtroProduto.Nome+"%", "%"+filtroProduto.Codigo+"%").Find(&resultQuery)
	} else {
		con.Where(" id = ?", filtroProduto.Id).Find(&resultQuery)
	}
	return
}

func InsertProduto(con *gorm.DB, produto *models.Produto) models.Produto {
	if err := con.Create(&produto).Error; err != nil {
		log.Println(err)
	}
	return *produto
}

func UpdateProduto(con *gorm.DB, produto *models.Produto) models.Produto {
	if err := con.Save(&produto); err != nil {
		log.Println(err)
	}
	return *produto
}

func DeleteProduto(con *gorm.DB, id int64) (produto models.Produto) {
	con.Where(" id = ?", id).Delete(&produto)
	return produto
}
