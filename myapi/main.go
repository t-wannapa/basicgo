package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "hi.")
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status: "active"},
}

func getTodosHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": todos,
	})
}

func getTodosByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	c.JSON(http.StatusOK, t)
}

func createTodosHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	id := len(todos)

	id = id + 1
	t.ID = id
	todos[t.ID] = &t

	c.JSON(http.StatusCreated, "created todo.")
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

	r.GET("/hello", helloHandler)
	r.GET("/todos", getTodosHandler)
	r.GET("/todos/:id", getTodosByIdHandler)
	r.POST("/todos", createTodosHandler)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":1234")
}
