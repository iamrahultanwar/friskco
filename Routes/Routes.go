package Routes

import (
	"github.com/iamrahultanwar/friskco/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userGrp := r.Group("/api/user")
	{
		userGrp.GET("/", Controllers.GetUsers)
		userGrp.POST("/", Controllers.CreateUser)

	}
	return r
}
