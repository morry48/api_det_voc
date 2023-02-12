package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	//"gorm.io/driver/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"nothing-behind.com/sample_gin/features/vocabulary/infra/postgres/model"
	"nothing-behind.com/sample_gin/features/vocabulary/infra/postgres/seed"
	"os"
)

type Database struct {
	*gorm.DB
}

var (
	db *gorm.DB
)

func New() (*Database, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Printf("環境変数が読み込み出来ませんでした: %v", err)
	}

	db, err = gorm.Open("postgres", "host="+os.Getenv("DB_HOST")+" port="+os.Getenv("DB_PORT")+" user="+os.Getenv("DB_USERNAME")+" dbname="+os.Getenv("DB_DATABASE")+" password="+os.Getenv("DB_PASSWORD")+" sslmode=disable")
	if err != nil {
		fmt.Printf("DB接続出来ませんでした: %v", err)
		panic(err)
	}
	autoMigration()

	return &Database{db}, nil
}

func autoMigration() {
	db.AutoMigrate(&model.Vocabulary{})
}

func Close() {
	if err := Database.Close; err != nil {
		panic(err)
	}
}

func InitForTest() (*Database, error) {
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
	return &Database{db}, nil
}
