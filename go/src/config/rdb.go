package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/joho/godotenv"
	"nothing-behind.com/sample_gin/features/vocabulary/infra/postgres/model"
	"nothing-behind.com/sample_gin/features/vocabulary/infra/postgres/seed"
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
	db, err = gorm.Open("postgres", "host="+os.Getenv("DB_HOST")+" port="+os.Getenv("DB_PORT")+" user="+os.Getenv("DB_USERNAME")+" dbname="+os.Getenv("DB_DATABASE")+" password="+os.Getenv("DB_PASSWORD")+" sslmode=disable")
	if err != nil {
		fmt.Printf("DB接続出来ませんでした: %v", err)
		panic(err)
	}
	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing postgres
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&model.Vocabulary{})
}

func InitForTest() {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		fmt.Printf("環境変数が読み込み出来ませんでした: %v", err)
	}

	db, err = gorm.Open("postgres", "host="+os.Getenv("DB_HOST")+" port="+os.Getenv("DB_PORT")+" user="+os.Getenv("DB_USERNAME")+" dbname="+os.Getenv("DB_DATABASE")+" password="+os.Getenv("DB_PASSWORD")+" sslmode=disable")
	if err != nil {
		fmt.Printf("DB接続出来ませんでした: %v", err)
		panic(err)
	}
	autoMigration()
	db.Exec("delete from vocabularies")
	db.Exec(seed.InitSql)
	if err != nil {
		fmt.Printf("データベース初期化失敗しました: %v", err)
	}
}
