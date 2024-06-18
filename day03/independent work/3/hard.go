/* добавление параметров:
- PUT /tasks/:id - обновить инф о задаче по ее ID
- DELETE /tasks/:id - удалить задачу по ее ID
- GET /tasks/:id - получить инф о задаче по ее ID
*/

package main

import (
		"github.com/gin-gonic/gin"
		"net/http"
		"github.com/google/uuid"
)

//структура задачи
type Task struct {
	UID uuid.UUID `json:"uid"`
	NewTask string `json:"newtask"`
}

//структура для хранения задач
type Task_Map struct {
	TaskMap map[uuid.UUID]Task `json:"taskmap"`
}

var task Task

//инициализация хранилища задач
var taskMap = Task_Map{TaskMap: make(map[uuid.UUID]Task)}

func main () {
	r := gin.Default()
	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, taskMap.TaskMap)
	})
	r.POST("/tasks", createTask)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)
	r.GET("/tasks/:id", getTask)
	r.Run(":8080")
}

func createTask(c *gin.Context) {


	if err := c.ShouldBindBodyWithJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
}
	// Генерация уникального идентификатора для задачи
	task.UID = uuid.New()

	idString := task.UID.String()


	// Добавление новой задачи в мапу
	taskMap.TaskMap[task.UID] = task

	// Возвращение успешного ответа с добавленной задачей
	c.JSON(http.StatusOK, gin.H{"message": "Added new task", "task": task, "id": task.UID})
}

func updateTask(c *gin.Context) {
	var uid uuid.UUID = c.Param("id")
	taskMap.TaskMap[uid] = task
	c.JSON(http.StatusOK, gin.H{"message": "Task updated", "id": id})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	id, ok := taskMap.task[uid]
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted", "id": id})
}

func getTask(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Task retrieved", "id": id})
}



