/* Расширьте API для управления списком задач (To-Do List). Добавьте следующие маршруты:
GET /tasks: Получить список задач.
POST /tasks: Создать новую задачу с использованием JSON-тела запроса. Проверьте, что обязательные поля (например, Title) присутствуют и имеют непустые значения
*/

package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Task struct {
	Title string `json:"title" validate:"notblank"`
	Info  string `json:"info" validate:"notblank"`
}

var list = []Task{}

var validate = validator.New()

// Пользовательская функция валидации для проверки на пустую строку
func notBlank(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

func main() {

	err := validate.RegisterValidation("notblank", notBlank)
	if err != nil {
		log.Fatalf("Error registering validation: %v", err)
	}
	
	r := gin.Default()
	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(200, list)
	})
	r.POST("/tasks", createTask)
	log.Println("Server starting on :8080")
	err = r.Run(":8080")
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

	if err := validate.Struct(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list = append(list, task)

	c.JSON(http.StatusOK, gin.H{"message": "Added new task", "task": task})

}
