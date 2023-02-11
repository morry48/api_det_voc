package postgres

import "github.com/jinzhu/gorm"

type Database struct {
	*gorm.DB
}

func New() (*Database, error) {
	var db *gorm.DB
	return &Database{db}, nil
}
