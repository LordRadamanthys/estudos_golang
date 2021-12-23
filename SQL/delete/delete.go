package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_credentials = "root:bebop268170@/cursogo"
)

func main() {
	db, _ := sql.Open("mysql", db_credentials)
	defer db.Close()
	stm, err := db.Prepare("DELETE from usuarios where id = ?")

	if err != nil {
		panic(err)
	}
	stm.Exec(201)

}
