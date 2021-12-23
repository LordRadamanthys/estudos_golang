package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_credentials = "root:bebop268170@/cursogo"
)

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para a func adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpe... ")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := createConexion()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}
	defer db.Close()

	var u Usuario

	//QueryRow traz apenas uma linha
	db.QueryRow("select * from usuarios where id =?", id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := createConexion()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}
	defer db.Close()

	var arrUsuarios []Usuario

	rows, _ := db.Query("select * from usuarios")
	defer rows.Close()

	for rows.Next() {
		var u Usuario
		rows.Scan(&u.ID, &u.Nome)
		arrUsuarios = append(arrUsuarios, u)
	}

	json, _ := json.Marshal(arrUsuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func createConexion() (*sql.DB, error) {
	db, err := sql.Open("mysql", db_credentials)

	if err != nil {
		return nil, err
	}
	return db, nil
}
