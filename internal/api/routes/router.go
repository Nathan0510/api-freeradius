package routes

import (
	"github.com/gin-gonic/gin"
	"api-freeradius/internal/api/controllers"
)

func SetupRouter(r *gin.Engine){

	{
		api := r.Group("/api")
		users := api.Group("/users")
		{
			users.POST("", controllers.CreateUser)
			users.GET("", controllers.GetAllUsers)
			users.GET("/:username", controllers.GetUser)
			users.DELETE("/:username", controllers.DeleteUser)
			users.PATCH("/:username", controllers.UpdateUser)
			users.DELETE("/:username/options/:optionname", controllers.DeleteUserOption)
		}
		
		nas := api.Group("/nas")
		{
			nas.POST("", controllers.CreateNas)
			nas.GET("", controllers.GetAllNas)
			nas.GET("/:nasname", controllers.GetNas)
			nas.DELETE("/:nasname", controllers.DeleteNas)
			nas.PATCH("/:nasname", controllers.UpdateNas)
		}

		groups := api.Group("/groups")
		{
			groups.POST("", controllers.CreateGroup)
			groups.GET("", controllers.GetAllGroup)
			groups.GET("/:groupname", controllers.GetGroup)
			groups.DELETE("/:groupname", controllers.DeleteGroup)
			groups.PATCH("/:groupname", controllers.UpdateGroup)
		}
		
		accounting := api.Group("/accounting")
		{
			accounting.GET("/:username", controllers.GetAccounting)
		}
	}
}