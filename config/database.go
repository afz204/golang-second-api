package config

import (
	"api1/entity"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

//SetupDatabase
func SetupDatabase() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Please check .env file!")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	con := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(con), &gorm.Config{})
	if err != nil {
		panic("Failed to create connection db!")
	}

	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}

func CloseConnection(db *gorm.DB)  {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close connection db!")
	}

	dbSql.Close();
}