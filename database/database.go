package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "root:1q2w3e4r@tcp(localhost:3307)/web_promotion_dev?charset=utf8mb4&parseTime=True&loc=Local"

	// เปิดการเชื่อมต่อฐานข้อมูล
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connected to MySQL successfully!")
	return db
}

func CloseDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database instance")
	}
	if err := sqlDB.Close(); err != nil {
		panic("Failed to close database connection")
	}
}
