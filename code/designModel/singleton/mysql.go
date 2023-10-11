package main

//import (
//	"sync"
//
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//var (
//	db   *gorm.DB
//	once sync.Once
//)
//
//func GetDB() *gorm.DB {
//	once.Do(func() {
//		// 初始化数据库连接池
//		db, err = gorm.Open(mysql.Open("dsn"), &gorm.Config{})
//		if err != nil {
//			panic(err)
//		}
//	})
//	return db
//}
