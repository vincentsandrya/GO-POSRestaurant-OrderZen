package models

import "time"

type User struct {
	UserId      int       `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"not null"`
	Email       string    `json:"email" gorm:"not null;unique"`
	Password    string    `json:"password"`
	RoleId      int       `json:"role_id"`
	Role        *Role     `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleId;references:RoleId"`
	CreatedDate time.Time `json:"created_date" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy   string    `json:"created_by"`
	UpdatedDate time.Time `json:"updated_date" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedBy   string    `json:"updated_by"`
}

type Role struct {
	RoleId int    `json:"role_id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name" gorm:"unique;not null"`
}
