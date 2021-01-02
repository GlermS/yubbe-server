package mysql

import (
	"database/sql"
)

type MySql struct {
	User     string
	Password string
	Schema   string
}

func (dbr *MySql) Open() (*sql.DB, error) {
	sqlDB, err := sql.Open("mysql", dbr.User+":"+dbr.Password+"@/"+dbr.Schema)
	if err != nil {
		panic(err.Error)
	}
	return sqlDB, err
}
