package login

import (
	"log"
	"yubbe-server/DataBase/DB"
	"yubbe-server/DataBase/interfaces"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type DBLoginVerifier struct {
	database interfaces.IDatabase
}

func NewDBLoginVerifier() *DBLoginVerifier {
	DBUser, ok := viper.Get("DB.USERNAME").(string)
	if !ok {
		log.Println("Erro carregando o username do banco de dados")
	}

	DBPassword, ok := viper.Get("DB.PASSWORD").(string)
	if !ok {
		log.Println("Erro carregando a senha do banco de dados")
	}

	DBSchema, ok := viper.Get("DB.NAME").(string)
	if !ok {
		log.Println("Erro carregando o nome do banco de dados")
	}

	db := &DB.Database{DBUser, DBPassword, DBSchema}
	return &DBLoginVerifier{db}
}

func (db *DBLoginVerifier) Check(user interfaces.ILoginUser) error {
	return db.database.Verify(user)
}
