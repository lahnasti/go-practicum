/*
Создайте RESTful API для управления списком задач (To-Do List). Реализуйте следующие маршруты:
GET /tasks: Получить список всех задач.
GET /tasks/:id: Получить информацию о задаче по ее ID.
POST /tasks: Создать новую задачу.
PUT /tasks/:id: Обновить информацию о задаче по ее ID.
DELETE /tasks/:id: Удалить задачу по ее ID.
Используйте структуры и слайсы для хранения информации о задачах. Каждая задача должна иметь уникальный идентификатор (ID), название (Title), описание (Description) и статус (Status), который может быть "Новая", "В процессе" или "Завершена".
*/

package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID    `json:"id"`
	Title       string `json:"title" validate:"notblank"`
	Description string `json:"description" validate:"notblank"`
	Status      string `json:"status" validate:"status"` 
}

// структура для хранения задач
type TaskMap struct {
	List map[string]Task `json:"list"`
}

var (
	taskMap = TaskMap{List: make(map[string]Task)}
	validate = validator.New()
	allowedStatus = []string{"New", "End", "In progress"}
	defaultStatus = "New"
)

func notBlank(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

// для проверки статуса
func validateStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	for _, s := range allowedStatus {
		if status == s {
			return true
		}
	}
	return false
}
func main() {

	err := validate.RegisterValidation("notblank", notBlank)
	if err != nil {
		log.Fatalf("Error registering validation: %v", err)
	}

	err = validate.RegisterValidation("status", validateStatus)
	if err != nil {
		log.Fatalf("Error registering status validation: %v", err)
	}

	r := gin.Default()
	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(200, taskMap.List)
	})
	r.POST("/tasks", createTasks)
	r.GET("/tasks/:id", getTasksId)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	log.Println("Server starting on :8080")

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func createTasks(c *gin.Context) {
	var task Task

	if err := c.ShouldBindBodyWithJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем дефолтное значение для Status, если оно не указано
	if task.Status == "" {
		task.Status = defaultStatus
	}

	if err := validate.Struct(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

		// Генерация уникального идентификатора для задачи
		task.ID = uuid.New()

		idString := task.ID.String()

	taskMap.List[idString] = task

	c.JSON(http.StatusOK, gin.H{"message": "Added new task", "task": task})

}

func getTasksId(c *gin.Context) {
	id := c.Param("id")
	task := taskMap.List[id]
	c.JSON(http.StatusOK, gin.H{"message": "Task retrieved", "task": task})

}

func updateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindBodyWithJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	taskMap.List[id] = task
	c.JSON(200, gin.H{"message": "Task updated", "id": id})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	delete(taskMap.List, id)
	c.JSON(200, gin.H{"message": "Task deleted", "id": id})
}
