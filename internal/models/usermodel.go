package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `gorm:"not null" json : "password" `
	Email    string `gorm:"unique" json : email`
}
