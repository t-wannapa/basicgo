package task

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://moftatjt:KGURjLDCcP9xjQMGEvO-13prL78g2HAA@arjuna.db.elephantsql.com:5432/moftatjt")
	if err != nil {
		log.Fatal(err)
	}
}

func GetTodoByIdHandler(c *gin.Context) {
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	row := stmt.QueryRow(id)
	t := &Todo{}
	err = row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, t)
}

func CreateTodosHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2) RETURNING id", t.Title, t.Status)
	err := row.Scan(&t.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, t)
}
func GetTodosHandler(c *gin.Context) {
	status := c.Query("status")
	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	todos := []Todo{}
	for rows.Next() {
		t := Todo{}
		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		todos = append(todos, t)
	}
	tt := []Todo{}
	for _, item := range todos {
		if status != "" {
			if item.Status == status {
				tt = append(tt, item)
			}
		} else {
			tt = append(tt, item)
		}
	}
	c.JSON(http.StatusOK, tt)
}

func UpdateTodosHandler(c *gin.Context) {
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	row := stmt.QueryRow(id)
	t := &Todo{}
	err = row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if err := c.ShouldBindJSON(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stmt, err = db.Prepare("UPDATE todos SET status=$2, title=$3 WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if _, err := stmt.Exec(id, t.Status, t.Title); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, t)
}

func DeleteTodosHandler(c *gin.Context) {
	id := c.Param("id")
	stmt, err := db.Prepare("DELETE FROM todos WHERE id = $1")
	if err != nil {
		log.Fatal("can't prepare delete statement", err)
	}
	if _, err := stmt.Exec(id); err != nil {
		log.Fatal("can't execute delete statment", err)
	}
	c.JSON(http.StatusOK, "deleted todo.")
}
