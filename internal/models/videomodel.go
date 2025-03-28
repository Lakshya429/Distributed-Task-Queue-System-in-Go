package models

import "time"

type Video struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Title       string    `gorm:"not null" json:"title" validate:"required,max=255"`
    FileName    string    `gorm:"unique" json:"file_name" validate:"required"`
    Description string    `gorm:"type:text" json:"description,omitempty"`
    UserID      uint      `gorm:"not null" json:"user_id"`
    User        *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
    
    // Processing metadata
    Status      string    `gorm:"default:'pending'" json:"status"`
    Duration    float64   `json:"duration,omitempty"`
    FileSize    int64     `json:"file_size,omitempty"` 
    // Timestamps
    CreatedAt   time.Time `gorm:"not null" json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at,omitempty"`
    ProcessedAt *time.Time `json:"processed_at,omitempty"`
}
