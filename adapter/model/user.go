package model

import "app/adapter/core"

type User struct {
	core.BaseModel
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
