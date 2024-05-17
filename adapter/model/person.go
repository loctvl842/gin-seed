package model

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
