/* RESTful API по управлению списком задач */

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	NewTask string `json:"newtask"`
}

type List struct {
	TaskList []Task `json:"tasklist"`
}

var list = List{TaskList: []Task{}}

// Переменная инициализируется как структура List с пустым срезом TaskList.
// Это позволяет сразу использовать list для добавления задач.

func main() {
	r := gin.Default()
	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, list)
	})
	r.POST("/tasks", createTask)
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func createTask(c *gin.Context) {

	var task Task

	if err := c.ShouldBindBodyWithJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Добавление новой задачи в список
	list.TaskList = append(list.TaskList, task)
	c.JSON(http.StatusOK, gin.H{"message": "Added new task", "task": task})
}
