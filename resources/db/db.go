package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
}

func GetUser(email string) (User, error) {
	var user User
	row := db.QueryRow("SELECT email, password FROM users WHERE email=$1", email)
	err := row.Scan(&user.Email, &user.Password)
	return user, err
}

type User struct {
	Email    string
	Password string
}
