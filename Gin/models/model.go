package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID    int      `json:"id" binding:"required,gte=1,lte=100"`
	Name  string  `json:"name" binding:"required,alpha"`
	Email string `json:"email" binding:"required,email"`
}

type User struct {
	ID    int     `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Email string   `json:"email" binding:"required,email"`
}