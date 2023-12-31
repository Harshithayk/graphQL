package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required"`
	HashPassword string `json:"hash_password" validate:"required"`
}

type Company struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	Salary   string `json:"salary" validate:"required"`
}
type Job struct {
	gorm.Model
	Company     Company `json:"-" gorm:"foreignKey:cid"`
	Cid         uint   `json:"cid"`
	Jobtitle    string `json:"title" validate:"required"`
	Salary      string `json:"salary" validate:"required"`
	Experience  string `json:"exp" validate:"required"`
}
