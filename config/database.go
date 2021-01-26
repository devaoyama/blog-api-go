package config

import (
	"blog-api/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Article{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
