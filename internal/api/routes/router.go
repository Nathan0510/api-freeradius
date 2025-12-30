package routes

import (
	"github.com/gin-gonic/gin"
	"api-freeradius/internal/api/controllers"
)

func SetupRouter(r *gin.Engine){

	{
		api := r.Group("/api")
		api.POST("/users", controllers.CreateUser)
		api.GET("/users", controllers.GetAllUsers)
		api.GET("/users/:username", controllers.GetUser)
		api.DELETE("/users/:username", controllers.DeleteUser)
		api.PATCH("/users/:username", controllers.UpdateUser)
		api.DELETE("/users/option/:username", controllers.DeleteUserOption)
		api.POST("/nas", controllers.CreateNas)
		api.GET("/nas", controllers.GetAllNas)
		api.GET("/nas/:username", controllers.GetNas)
		api.DELETE("/nas/:username", controllers.DeleteNas)
		api.PATCH("/nas/:username", controllers.UpdateNas)
	}
}