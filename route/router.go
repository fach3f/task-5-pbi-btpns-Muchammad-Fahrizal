// router/router.go

package router

import (
	"tugasakhir/controllers"
	"tugasakhir/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users/register", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)
	r.PUT("/users/:userId", middleware.Auth, controllers.UpdateUser)
	r.DELETE("/users/:userId", middleware.Auth, controllers.DeleteUser)
	r.POST("/photos", middleware.Auth, controllers.CreatePhoto)
	r.GET("/photos", middleware.Auth, controllers.ShowPhoto)
	r.PUT("/photos/:photoId", middleware.Auth, controllers.UpdatePhoto)
	r.DELETE("/photos/:photoId", middleware.Auth, controllers.DeletePhoto)
	
	return r
}
