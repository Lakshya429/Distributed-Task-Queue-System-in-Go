package repository 

import (
	"github.com/Lakshya429/distributed-task-queue/pkg/database"
	"github.com/Lakshya429/distributed-task-queue/internal/models"
)

var DB = database.GetDB()
func CreateVideo(video * models.Video)  error {
	return DB.Create(video).Error
}

func CreateUser(user * models.User) error {
	return DB.Create(user).Error
}
func GetAllVideos() ([] models.Video , error) {
	var video [] models.Video
	err := DB.Find(&video).Error
	return video, err
}

func GetVideoByUser(userID string) ([] models.Video , error) {
	var video [] models.Video
	err := DB.Where("user_id = ?" , userID).Find(&video).Error
	return video , err
}