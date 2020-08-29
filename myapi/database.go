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

func updateTodo() {
	db, err := sql.Open("postgres", "postgres://moftatjt:KGURjLDCcP9xjQMGEvO-13prL78g2HAA@arjuna.db.elephantsql.com:5432/moftatjt")
	if err != nil {
		log.Fatal("connect to database error", err)
	}
	defer db.Close()
	
	stmt, err := db.Prepare("UPDATE todos SET status=$2 WHERE id=$1")
	if err != nil {
		log.Fatal("can't prepare statement update", err)
	}
	
	if _, err := stmt.Exec(1, "inactive"); err != nil {
		log.Fatal("error execute update", err)
	}
	
	fmt.Print("error execute update", err)
}

func queryAllTodo() {
	db, err := sql.Open("postgres", "postgres://moftatjt:KGURjLDCcP9xjQMGEvO-13prL78g2HAA@arjuna.db.elephantsql.com:5432/moftatjt")
	if err != nil {
		log.Fatal("connect to database error", err)
	}

	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}

	for rows.Next() {
		var id int
		var title, status string
		err := rows.Scan(&id, &title, &status)
		if err != nil {
			log.Fatal("can't scan row into variable", err)
		}
		fmt.Println(id, title, status)
	}

	fmt.Println("query all todos success")
}

func main() {
	//insertTodo()
	//getTodo()
	updateTodo()
}
