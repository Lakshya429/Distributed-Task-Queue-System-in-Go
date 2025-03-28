package repository

import (
	"errors"
	"log"
	"github.com/Lakshya429/distributed-task-queue/internal/models"
	"github.com/Lakshya429/distributed-task-queue/pkg/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	DB = database.GetDB()
}

func CreateVideo(video *models.Video) error {
	return DB.Create(video).Error
}

func CreateUser(user *models.User) error {
	if user == nil || DB == nil{
		return errors.New("user is nil")
	}
	log.Println(user)
	return DB.Create(user).Error
}
func GetAllVideos() ([]models.Video, error) {
	var video []models.Video
	err := DB.Find(&video).Error
	return video, err
}

func GetVideoByUser(userID string) ([]models.Video, error) {
	var video []models.Video
	err := DB.Where("user_id = ?", userID).Find(&video).Error
	return video, err
}

func GetUserbyUserName(username string) (*models.User, error) {
	var user models.User
	err := DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
