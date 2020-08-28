package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func insertTodo() {
	db, err := sql.Open("postgres", "postgres://moftatjt:KGURjLDCcP9xjQMGEvO-13prL78g2HAA@arjuna.db.elephantsql.com:5432/moftatjt")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id", "buy bmw", "active")
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	}
	fmt.Println("insert todo success id : ", id)

	defer db.Close()
}

func getTodo() {
	db, err := sql.Open("postgres", "postgres://moftatjt:KGURjLDCcP9xjQMGEvO-13prL78g2HAA@arjuna.db.elephantsql.com:5432/moftatjt")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)
	var id int
	var title, status string
	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	fmt.Println("one row", id, title, status)

	defer db.Close()
}

func main() {
	//insertTodo()
	//getTodo()
	db, err := sql.Open("postgres", "postgres://moftatjt:KGURjLDCcP9xjQMGEvO-13prL78g2HAA@arjuna.db.elephantsql.com:5432/moftatjt")
	if err != nil {
		log.Fatal("connect to database error", err)
	}

	stmt, err := db.Prepare("UPDATE todos Set")

	defer db.Close()
}
