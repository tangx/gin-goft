package adaptors

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormAdaptor struct {
	*gorm.DB
}

func NewGormAdaptor() *GormAdaptor {
	dsn := "root:Mysql12345@tcp(127.0.0.1:3306)/goftdemo?charset=utf8mb4&parseTime=True&loc=Local"
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &GormAdaptor{
		DB: gormdb,
	}
}
