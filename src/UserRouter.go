package src

import (
	"golangAPI/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("users")
	user.GET("/", service.FindAllUsers)
	user.GET("/search", service.FindByUserId)
	user.POST("/", service.PostUser)
	user.DELETE("/search", service.DeleteUser)
	user.PUT("/search", service.PutUser)
}
