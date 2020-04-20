package database

import (
	"fmt"
	"os"

	"github.com/Manah6Eugenio/faturacao/utils"
)

//variáveis de ambiente com as configurações de conexão com o banco de dados
const (
	DBHOST    = "DBHOST"
	DBPORT    = "DBPORT"
	DBUSER    = "DBUSER"
	DBPASS    = "DBPASS"
	DBNAME    = "DBNAME"
	DBSSLMODE = "DBSSLMODE"
)

//função que retorna o driver de configuração do banco de dados
func driverConfig() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv(DBHOST), utils.ParseInt(os.Getenv(DBPORT)), os.Getenv(DBUSER), os.Getenv(DBPASS), os.Getenv(DBNAME), os.Getenv(DBSSLMODE))
}
