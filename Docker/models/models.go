package models

import "gorm.io/gorm"

type Measurement struct {
	gorm.Model
	Quantity string `json:"quantity" gorm:"text;not null;default:null"`
	Value    string `json:"value" gorm:"text;not null;default:null"`
}
