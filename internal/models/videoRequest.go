package models 

type VideoRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	UserID uint `gorm : "not null"json:"user_id"validate:"required"`
	FileName string `json:"file_name"`
}