//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"nothing-behind.com/sample_gin/packages/vocabulary/infra/postgres/repository"
	"nothing-behind.com/sample_gin/packages/vocabulary/usecase"
)

type App struct {
	rdb *posgre.Database

	notificationApp *notification.App
}

func InitVocabularyList(db *gorm.DB) usecase.ListVocabularies {
	wire.Build(
		usecase.NewListCategories,
		repository.NewVocabularyRepository,
	)
	return nil
}
