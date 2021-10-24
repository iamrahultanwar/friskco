package Models

import (
	"github.com/iamrahultanwar/friskco/Config"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name    string `json:"name"`
	DriveID int    `json:"driveId"`
	Path    string `json:"path"`
	UserID  int    `json:"userId"`
}

func GetUserFilesByDriveId(files *[]File, driveId int) {

}

func SaveFileData(file *File) error {
	if err := Config.DB.Create(file).Error; err != nil {
		return err
	}
	return nil
}
