package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamrahultanwar/friskco/Models"
)

func GetAllUserDrives(c *gin.Context) {
	var drives []*Models.Drive
	userId := c.GetFloat64("userId")
	drives, _, err := Models.GetAllUserDrives(userId, drives)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.AbortWithStatusJSON(http.StatusOK, drives)
}

func CreateUserDrive(c *gin.Context) {

	userId := c.GetFloat64("userId")

	var drive *Models.Drive
	c.BindJSON(&drive)
	drive.UserID = int(userId)

	type createUserDriveMessage struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}

	if err := Models.CreateUserDrive(drive); err != nil {
		message := createUserDriveMessage{
			Status:  false,
			Message: "Error in creating drive",
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, message)
	}
	message := createUserDriveMessage{
		Status:  true,
		Message: "Drive created",
	}
	c.AbortWithStatusJSON(http.StatusOK, message)

}

func GetAllUserFiles(c *gin.Context) {
	files := []Models.File{}
	driveId := c.Param("driveId")

	if x, err := strconv.ParseInt(driveId, 10, 64); err == nil {
		if err := Models.GetAllDriveFiles(&files, int(x)); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false})
		}
		c.AbortWithStatusJSON(http.StatusOK, files)
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Something went wrong"})
}
