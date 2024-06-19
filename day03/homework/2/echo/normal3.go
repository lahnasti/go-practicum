/* Echo — это высокопроизводительный, минималистичный HTTP-фреймворк для Go,
который предлагает удобный API для построения RESTful приложений.

Преимущества:

- Высокая производительность.
- Хорошо поддерживает Middleware.
- Простота в использовании и богатая документация.

*/

package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Task struct {
	NewTask string `json:"newtask"`
}

var tasks = []Task{}

func main() {
	e := echo.New()

	e.GET("/tasks", getTasks)
	e.POST("/tasks", createTask)

	log.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", e)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func getTasks(c echo.Context) error {

	return c.JSON(http.StatusOK, tasks)
}

func createTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	tasks = append(tasks, task)
	return c.JSON(http.StatusCreated, task)
}
