package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/joho/godotenv"
	"nothing-behind.com/sample_gin/entity"
	"os"
)

var (
	db *gorm.DB
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("環境変数が読み込み出来ませんでした: %v", err)
	}
	os.Getenv("DB_HOST")

	db, err = gorm.Open("postgres", "host="+os.Getenv("DB_HOST")+" port="+os.Getenv("DB_PORT")+" user="+os.Getenv("DB_USERNAME")+" dbname="+os.Getenv("DB_DATABASE")+" password="+os.Getenv("DB_PASSWORD")+" sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.Vocabulary{})
}
