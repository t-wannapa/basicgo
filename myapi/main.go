package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/t-wannapa/myapi/middleware"
	"github.com/t-wannapa/myapi/task"
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

func setupRouter() *gin.Engine {
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.Auth)
	apiV1.GET("/todos", task.GetTodosHandler)
	apiV1.GET("/todos/:id", task.GetTodoByIdHandler)
	apiV1.POST("/todos", task.CreateTodosHandler)
	apiV1.PUT("/todos/:id", task.UpdateTodosHandler)
	apiV1.DELETE("/todos/:id", task.DeleteTodosHandler)
	return r
}
func main() {
	r := setupRouter()
	r.Run(":1234")
}
