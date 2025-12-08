package routes

import (
	"github.com/gin-gonic/gin"
	"api-freeradius/internal/api/controllers"
)

func SetupRouter(r *gin.Engine){

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:username", controllers.GetUser)
	r.POST("/nas", controllers.CreateNas)
	r.GET("/nas", controllers.GetAllNas)
	r.GET("/nas/:username", controllers.GetNas)
	r.DELETE("/nas/:username", controllers.DeleteNas)
	r.PUT("/nas/:username", controllers.UpdateNas)

}