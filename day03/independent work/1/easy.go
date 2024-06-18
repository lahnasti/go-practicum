/* РЕАЛИЗАЦИЯ ПРОСТОГО RESTful API с использованием Gin */

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	r.Run(":8080")
}

// go run easy.go

// контекст используется для доступа к информации о запросе 
// и для формирования ответа.