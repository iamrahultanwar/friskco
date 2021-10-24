package Routes

import (
	"net/http"

	"github.com/iamrahultanwar/friskco/Controllers"
	"github.com/iamrahultanwar/friskco/Models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	userGrp := r.Group("/api/user")
	{
		userGrp.GET("/me", Models.Authorized(), Controllers.GetCurrentUser)
		userGrp.POST("/register", Controllers.RegisterUser)
		userGrp.POST("/login", Controllers.LoginUser)

	}

	driveGrp := r.Group("/api/drive", Models.Authorized())
	{
		driveGrp.GET("/", Controllers.GetAllUserDrives)
		driveGrp.POST("/create", Controllers.CreateUserDrive)
		driveGrp.GET("/files/:driveId", Controllers.GetAllUserFiles)

	}

	fileGrp := r.Group("/api/file")
	{
		fileGrp.POST("/upload", Controllers.UploadFile)
		fileGrp.StaticFS("/get/public", http.Dir("public"))

	}

	return r
}
