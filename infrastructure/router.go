package infrastructure

import (
	"github.com/gin-gonic/gin"
	"modul1/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHandler())

	router.POST("/signup", func(c *gin.Context) { userController.Create(c) })
	router.POST("/signin", func(c *gin.Context) { userController.Compare(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}
