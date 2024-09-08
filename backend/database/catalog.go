package database

import "gorm.io/gorm"

type catalog struct {
	gorm.Model            // 内嵌gorm.Model,包含ID,CreatedAt,UpdatedAt,DeletedAt
	path       string     `gorm:"not null;size:100"`
	ParentID   uint       `gorm:"default:null"`
	Parent     *catalog   `gorm:"foreignKey:ParentID"`
	Children   []*catalog `gorm:"foreignKey:ParentID"`
	IsFile     bool       `gorm:"not null"`
	IsShow     bool       `gorm:"not null;default:true"`
}
