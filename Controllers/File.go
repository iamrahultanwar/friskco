package Controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/iamrahultanwar/friskco/Models"
)

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	driveId := c.PostForm("driveId")

	fileStruct := Models.File{}
	if x, err := strconv.ParseInt(driveId, 10, 64); err == nil {
		fileStruct.DriveID = int(x)
	}
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	fileStruct.Name = filename
	fileStruct.UserID = 1
	fSplit := strings.Split(fileStruct.Name, ".")
	slug := slug.Make(fSplit[0])
	filename = slug + "." + fSplit[1]
	fileStruct.Path = "public/" + filename

	out, err := os.Create(fileStruct.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	if err := Models.SaveFileData(&fileStruct); err != nil {
		os.Remove(fileStruct.Path)
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error in uploading"})

	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "File Uploaded", "path": fileStruct.Path})
}
