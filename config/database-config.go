package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/andiahmads/go-api/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SETUP DATABASE, CREATE CONNECTION
func SetupDatabaseConnection() *gorm.DB {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	fmt.Println(basepath)
	err := godotenv.Load(fmt.Sprintf("%s/../.env", basepath))

	if err != nil {
		panic("Failed, env not found!")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	//nanti kita isi modelnya di sini
	// db.AutoMigrate()
	db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Categories{})
	return db
}

//CloaseDatabaseConnection method is closing a connection between your app and db
//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}

//
