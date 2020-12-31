package register

import (
	"yubbe-server/DataBase/DB"
	"yubbe-server/DataBase/interfaces"

	_ "github.com/go-sql-driver/mysql"
)

type DBRegister struct {
	database interfaces.IDatabase
}

func NewDBRegister() *DBRegister {
	db := &DB.Database{"root", "#Santos075", "dbname"}
	return &DBRegister{db}
}

func (dbr *DBRegister) Write(user interfaces.IRegisterUser) error {
	return dbr.database.Register(user)
}
