package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	wspaces := getAllWorkspaces()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":    "Home Page",
			"workload": wspaces,
		},
	)

}
func getWorkSpace(c *gin.Context) {
	// Проверим валидность ID
	if workSpaceID, err := strconv.Atoi(c.Param("workspace_id")); err == nil {
		// Проверим существование топика
		if workSpace, err := getWorkSpaceByID(workSpaceID); err == nil {
			// Вызовем метод HTML из Контекста Gin для обработки шаблона
			c.HTML(
				// Зададим HTTP статус 200 (OK)
				http.StatusOK,
				// Используем шаблон index.html
				"workspace.html",
				// Передадим данные в шаблон
				gin.H{
					"name":     workSpace.Name,
					"workload": workSpace,
				},
			)

		} else {
			// Если топика нет, прервём с ошибкой
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// При некорректном ID в URL, прервём с ошибкой
		c.AbortWithStatus(http.StatusNotFound)
	}
}
