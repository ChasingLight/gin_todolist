package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局变量
var (
	DB *gorm.DB
)

// InitMysql 建立mysql连接
func InitMysql() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
