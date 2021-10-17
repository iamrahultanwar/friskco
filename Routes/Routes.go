package Routes

import (
	"github.com/iamrahultanwar/friskco/Controllers"
	"github.com/iamrahultanwar/friskco/Models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	userGrp := r.Group("/api/user")
	{
		userGrp.GET("/", Models.Authorized(), Controllers.GetUsers)
		userGrp.POST("/register", Controllers.RegisterUser)
		userGrp.POST("/login", Controllers.LoginUser)

	}

	return r
}
