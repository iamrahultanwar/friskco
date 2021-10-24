package Models

import (
	"github.com/iamrahultanwar/friskco/Config"
	"gorm.io/gorm"
)

type Drive struct {
	gorm.Model
	Name   string `json:"name"`
	UserID int    `json:"userId"`
}

func GetAllUserDrives(userID float64, drives []*Drive) ([]*Drive, int, error) {
	result := Config.DB.Where("user_id = ?", int(userID)).Find(&drives)
	return drives, int(result.RowsAffected), result.Error
}

func CreateUserDrive(drive *Drive) error {
	if err := Config.DB.Create(drive).Error; err != nil {
		return err
	}
	return nil
}

func GetAllDriveFiles(files *[]File, driveId int) error {

	if err := Config.DB.Where("drive_id = ?", driveId).Find(&files).Error; err != nil {
		return err
	}
	return nil

}
