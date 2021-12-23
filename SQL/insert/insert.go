package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_credentials = "root:bebop268170@/cursogo"
)

func main() {
	db, err := sql.Open("mysql", db_credentials)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")

	// stmt.Exec("Maria")
	// stmt.Exec("Mateus")

	res, _ := stmt.Exec("Lucia")
	id, _ := res.LastInsertId()
	fmt.Println(id)
}
