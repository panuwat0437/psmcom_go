package db

import (
	"main/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDB - call this method to get db
func GetDB() *gorm.DB {
	return db
}

// SetupDB - setup dabase for sharing to all api
func SetupDB() {

	// sqlite "gorm.io/driver/sqlite"
	// database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	// mysql "gorm.io/driver/mysql"
	dsn := "prinpt_dcscore:eQ6hmoOJVsir@tcp(156.67.219.130)/prinpt_dcscore?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// postgresql 	"gorm.io/driver/postgres"
	// dsn := "user=postgres password=12341234 dbname=cmgostock port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	// database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// database.AutoMigrate(&model.User{})
	// database.AutoMigrate(&model.Product{})
	// database.AutoMigrate(&model.Transaction{})
	// database.AutoMigrate(&model.Score{})
	// database.AutoMigrate(&model.Month{})

	db = database
}
