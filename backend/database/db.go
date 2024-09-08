package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB // 全局数据库连接实例

func InitDB() {
	dsn := "root:@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=UTC"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
}
