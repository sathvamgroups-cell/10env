package routes

import (
	"10env/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, controller *controllers.UserController) {
	router.POST("/users", controller.CreateUser)
	router.GET("/users", controller.GetUsers)
	router.GET("/users/:id", controller.GetUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
}
