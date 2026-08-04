package routes

import (
	"github.com/gin-gonic/gin"

	"example.com/handlers"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/projects", handlers.CreateProject)
	r.GET("/projects", handlers.GETHome)

	return r
}
