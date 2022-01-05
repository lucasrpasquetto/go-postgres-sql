package database\conf

import (
	"database/sql"
	"fmt"
)

const Driver = "postgres"
const Host = "127.0.0.1"
const Port = "5432"
const User = "go"
const Password = "go"
const Database = "go"

var SourceName = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Database)
var stantment = *sql.DB

func connectDb() *sql.DB {
	stantment, err = sql.Open(Driver, SourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Conected!")
	}
	return stantment
}
