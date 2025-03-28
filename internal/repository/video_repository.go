package repository 

import (
	"github.com/Lakshya429/distributed-task-queue/pkg/database"
	"github.com/Lakshya429/distributed-task-queue/internal/models"
)

var DB = database.GetDB()
func CreateVideo(video * models.Video)  error {
	return DB.Create(video).Error
}

func GetAllVideos() ([] models.Video , error) {
	var video [] models.Video

	err = DB.
}