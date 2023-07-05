package db

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var client *gorm.DB

func init() {
	dsn := "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, "root", "123456", "10.10.10.43", "3306", "monitor")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	client = db
}

func GetReadDB(ctx context.Context) *gorm.DB {
	return client.WithContext(ctx)
}

func GetWriteDB(ctx context.Context) *gorm.DB {
	return client.WithContext(ctx)
}
