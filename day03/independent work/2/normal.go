/* RESTful API по управлению списком задач */

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Task struct {
	NewTask string `json:"newtask"`
}

type List struct {
	TaskList []Task `json:"tasklist"`
}

 // Переменная инициализируется как структура List с пустым срезом TaskList. 
// Это позволяет сразу использовать list для добавления задач.
var list = List{TaskList: []Task{}}

func main() {
	r := gin.Default()
	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, list)
	})
	r.POST("/tasks", createTask)
	r.Run(":8080")
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
