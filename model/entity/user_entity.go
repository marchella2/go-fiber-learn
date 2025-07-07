package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	//gorm.Model        // default dari GORM (value nya ada id, created_at, updated_at, deleted_at
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
