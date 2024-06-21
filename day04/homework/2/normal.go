/*
Расширьте предыдущий API для управления списком пользователей. Создайте маршруты:
GET /users: Получить список всех пользователей.
GET /users/:id: Получить информацию о пользователе по его ID.
POST /users: Создать нового пользователя.
PUT /users/:id: Обновить информацию о пользователе по его ID.
DELETE /users/:id: Удалить пользователя по его ID.
Каждый пользователь должен иметь уникальный идентификатор (ID), имя (Name), адрес электронной почты (Email) и пароль (Password).
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

type User struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name" validate:"notblank"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"notblank"`
}

// структура для хранения списка пользователей
type UserMap struct {
	ListUser map[string]User `json:"listuser"`
}

var (
	taskMap = TaskMap{List: make(map[string]Task)}
	validate = validator.New()
	allowedStatus = []string{"New", "End", "In progress"}
	defaultStatus = "New"

	userMap = UserMap{ListUser: make(map[string]User)}
)

// проверка на пустую строку в json
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

	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.GET("/", getTasks)
		taskRoutes.POST("/", createTasks)

		taskRoutes.GET("/:id", getTasksId)
		taskRoutes.PUT("/:id", updateTask)
		taskRoutes.DELETE("/:id", deleteTask)
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", getUsers)
		userRoutes.POST("/", createUser)
		
		userRoutes.GET("/:id", getUserId)
		userRoutes.PUT("/:id", updateUser)
		userRoutes.DELETE("/:id", deleteUser)
	}


	log.Println("Server starting on :8080")

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func getTasks(c *gin.Context) {
	c.JSON(200, taskMap.List)
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

func getUsers(c *gin.Context) {
		c.JSON(200, userMap.ListUser)
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uuid.New()
	idString := user.ID.String()

	userMap.ListUser[idString] = user
	c.JSON(200, gin.H{"message": "Added new user", "user": user})
}

func getUserId(c *gin.Context) {
	id := c.Param("id")
	user := userMap.ListUser[id]
	c.JSON(200, gin.H{"message": "User retrieved", "user": user})
}

func updateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	userMap.ListUser[id] = user
	c.JSON(200, gin.H{"message": "User updated", "user": user})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	delete(userMap.ListUser, id)
	c.JSON(200, gin.H{"message": "User deleted", "id": id})
}