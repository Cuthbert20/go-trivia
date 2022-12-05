package models

import "gorm.io/gorm"

type fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null`
	Answer   string `json:"answer" gorm:"text;not null;default:null`
}
