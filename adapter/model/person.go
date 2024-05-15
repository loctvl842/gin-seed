package model

type Person struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
