/*
Создайте RESTful API для управления списком книг (Book List) с использованием пакета Gin. Реализуйте следующие маршруты:
GET /books: Получить список всех книг.
POST /books: Создать новую книгу с использованием JSON-тела запроса. Проверьте, что обязательные поля (например, Title, Author) присутствуют в запросе и имеют непустые значения.
*/

package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Book struct {
	Name   string `json:"name" validate:"notblank"`
	Author string `json:"author" validate:"notblank"`
}

var list = []Book{}

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
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, list)
	})
	r.POST("/books", createBook)
	log.Println("Server starting on :8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func createBook(c *gin.Context) {
	var book Book

	if err := c.ShouldBindBodyWithJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list = append(list, book)
	c.JSON(http.StatusOK, gin.H{"message": "Added new book", "book": book})

}
