package database

import "gorm.io/gorm"

type suffix struct {
	gorm.Model
	Name   string `gorm:"not null;size:8"`
	IsShow bool   `gorm:"not null;default:true"`
}
