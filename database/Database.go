package database

import (
	"fmt"
	"os"

	"github.com/Manah6Eugenio/faturacao/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//função que retorna o driver de configuração do banco de dados
func driverConfig() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DBHOST"), utils.ParseInt(os.Getenv("DBPORT")), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPASS"))
}

func Connection() *gorm.DB {
	db, err := gorm.Open("postgres", driverConfig())
	if err != nil {
		panic(err.Error())
	}
	db.SingularTable(true)
	return db
}
