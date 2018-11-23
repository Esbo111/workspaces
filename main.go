package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Создаем роутер
	router := gin.Default()
	//  Загружаем шаблоны
	router.LoadHTMLGlob("templates/*")

	// инициализируем роуты
	router.GET("/", func(c *gin.Context) {
		wspaces := getAllWorkspaces()

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":    "Home Page",
				"workload": wspaces,
			},
		)

	})

	router.GET("/workspace/view/:workspace_id", getWorkSpace)

	// запуск приложения
	router.Run()

}
