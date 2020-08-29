package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// var db *sql.DB

// func init() {
// 	var err error
// 	db, err := sql.Open("postgres", "postgres://moftatjt:KGURjLDCcP9xjQMGEvO-13prL78g2HAA@arjuna.db.elephantsql.com:5432/moftatjt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "hi.")
}

// var todos = map[int]*Todo{
// 	1: &Todo{ID: 1, Title: "pay phone bills", Status: "active"},
// }

func getTodosHandler(c *gin.Context) {
	status := c.Query("status")

	db, err := sql.Open("postgres", "postgres://moftatjt:KGURjLDCcP9xjQMGEvO-13prL78g2HAA@arjuna.db.elephantsql.com:5432/moftatjt")
	if err != nil {
		log.Fatal("connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}

	todos := []Todo{}
	for rows.Next() {
		var id int
		var title string
		var status string

		err := rows.Scan(&id, &title, &status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		t := Todo{
			ID:     id,
			Title:  title,
			Status: status,
		}
		todos = append(todos, t)
	}

	tt := []Todo{}
	for _, item := range todos {
		if item.Status == status {
			tt = append(tt, item)
		}
	}
	c.JSON(http.StatusOK, tt)
}

func updateTodosHandler(c *gin.Context) {
	id := c.Param("id")
	db, err := sql.Open("postgres", "postgres://moftatjt:KGURjLDCcP9xjQMGEvO-13prL78g2HAA@arjuna.db.elephantsql.com:5432/moftatjt")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

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

	c.JSON(http.StatusCreated, t)
}

func getTodosByIdHandler(c *gin.Context) {
	id := c.Param("id")
	db, err := sql.Open("postgres", "postgres://moftatjt:KGURjLDCcP9xjQMGEvO-13prL78g2HAA@arjuna.db.elephantsql.com:5432/moftatjt")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

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
	// 	id, err := strconv.Atoi(c.Param("id"))
}

func createTodosHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	db, err := sql.Open("postgres", "postgres://moftatjt:KGURjLDCcP9xjQMGEvO-13prL78g2HAA@arjuna.db.elephantsql.com:5432/moftatjt")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id", t.Title, t.Status)
	err = row.Scan(&t.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, t)
}

func authMiddleware(c *gin.Context) {
	log.Println("start middleware")
	authKey := c.GetHeader("Authorization")
	if authKey != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		c.Abort()
		return
	}
	c.Next()
	log.Println("end middleware")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(authMiddleware)
	apiV1 := r.Group("/api/v1")
	// r.GET("/hello", helloHandler)
	apiV1.GET("/todos", getTodosHandler)
	apiV1.GET("/todos/:id", getTodosByIdHandler)
	apiV1.POST("/todos", createTodosHandler)
	apiV1.PUT("/todos/:id", updateTodosHandler)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":1234")
}
