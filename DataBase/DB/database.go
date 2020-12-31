package DB

import (
	"YubbeServer/yubbe-server/DataBase/interfaces"
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Database struct {
	//	*mysql.MySql
	User     string
	Password string
	Schema   string
}

func (database *Database) GetDatabaseData() (string, string) {
	return database.User, database.Schema
}

func (database *Database) Register(user interfaces.IRegisterUser) error {
	db, _ := database.Open()
	defer db.Close()

	_, err := db.Query("INSERT INTO " + database.Schema + ".users (name, email, password) VALUES ('" + user.GetName() + "','" + user.GetEmail() + "','" + string(user.GetPassword()) + "')")
	//defer insert.Close()

	return err
}

func (database *Database) Verify(user interfaces.ILoginUser) error {
	db, err := database.Open()
	defer db.Close()

	rows, err := db.Query("SELECt email, password FROM " + database.Schema + ".users WHERE email = '" + user.GetEmail() + "'")
	//defer rows.Close()
	if err != nil {
		log.Println("Email não encontrado")
		return err
	}

	err = database.checkPassword(rows, user.GetPassword())

	return err
}

func (database *Database) checkPassword(rows *sql.Rows, password string) error {
	var (
		email string
		key   string
	)
	ok := errors.New("Error desconhecido no checkPassword")
	for rows.Next() {
		err := rows.Scan(&email, &key)
		if err != nil {
			log.Fatal(err)
		}
		ok = bcrypt.CompareHashAndPassword([]byte(key), []byte(password))
		if ok != nil {
			ok = errors.New("Senha incorreta")
		}
	}
	return ok
}

func (dbr *Database) Open() (*sql.DB, error) {
	sqlDB, err := sql.Open("mysql", dbr.User+":"+dbr.Password+"@/"+dbr.Schema)
	if err != nil {
		panic(err.Error)
	}
	return sqlDB, err
}
