/* добавление параметров:
- PUT /tasks/:id - обновить инф о задаче по ее ID
- DELETE /tasks/:id - удалить задачу по ее ID
- GET /tasks/:id - получить инф о задаче по ее ID
*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// структура задачи
type Task struct {
	UID     uuid.UUID `json:"uid"`
	NewTask string    `json:"newtask"`
}

// структура для хранения задач
type TaskMap struct {
	List map[string]Task `json:"list"`
}

// инициализация хранилища задач
var taskMap = TaskMap{List: make(map[string]Task)}

func main() {
	r := gin.Default()
	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, taskMap.List)
	})
	r.POST("/tasks", createTask)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)
	r.GET("/tasks/:id", getTask)
	r.Run(":8080")
}

func createTask(c *gin.Context) {

	var task Task
	if err := c.ShouldBindBodyWithJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Генерация уникального идентификатора для задачи
	task.UID = uuid.New()

	idString := task.UID.String()

	// Добавление новой задачи в мапу
	taskMap.List[idString] = task

	// Возвращение успешного ответа с добавленной задачей
	c.JSON(http.StatusOK, gin.H{"message": "Added new task", "task": task})
}

func updateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindBodyWithJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid := c.Param("id")
	taskMap.List[uid] = task
	c.JSON(http.StatusOK, gin.H{"message": "Task updated", "id": uid})
}

func deleteTask(c *gin.Context) {
	uid := c.Param("id")
	delete(taskMap.List, uid)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted", "id": uid})
}

func getTask(c *gin.Context) {

	uid := c.Param("id")
	task := taskMap.List[uid]
	c.JSON(http.StatusOK, gin.H{"message": "Task retrieved", "task": task})
}
