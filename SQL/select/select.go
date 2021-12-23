package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_credentials = "root:bebop268170@/cursogo"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, _ := sql.Open("mysql", db_credentials)

	defer db.Close()

	rows, _ := db.Query("select * from usuarios where id > ?", 0)
	defer rows.Close()

	var arrUser []usuario
	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		arrUser = append(arrUser, u)
		// fmt.Println(u)
	}
	fmt.Println(arrUser)
}
