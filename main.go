package main

import (
	"database/sql"
	"fmt"
	"strconv"
"github.com/lucasrpasquetto/go-postgres-sql/database/conf/db"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID        int
	Title     string
	Is_closed bool
}

var conn *sql.DB
var err error

func throwErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	fmt.Printf("Iniciando...")

	conn = 

	defer conn.Close()

	sqlStatment, err := conn.Query("select * from todo")
	throwErr(err)

	for sqlStatment.Next() {
		var todo Todo

		err = sqlStatment.Scan(&todo.ID, &todo.Title, &todo.Is_closed)
		throwErr(err)

		fmt.Printf("%d\t%s\t%s \n", todo.ID, todo.Title, strconv.FormatBool(todo.Is_closed))
	}
}
