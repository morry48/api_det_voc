package repository

import (
	"nothing-behind.com/sample_gin/config"
	"nothing-behind.com/sample_gin/packages/vocabulary/domain/entity"
	"nothing-behind.com/sample_gin/packages/vocabulary/domain/interface_repository"
	"nothing-behind.com/sample_gin/packages/vocabulary/infra/postgres"
	"nothing-behind.com/sample_gin/packages/vocabulary/infra/postgres/model"
)

func NewVocabularyRepository(db *postgres.Database) repository.VocabularyRepository {
	return &VocabularyRepositoryImpl{db: *db}
}

type VocabularyRepositoryImpl struct {
	db postgres.Database
}

func (v VocabularyRepositoryImpl) SelectByOption(option *entity.QueryOption) (*[]entity.Vocabulary, error) {

	detDb := config.GetDB()

	var result []entity.Vocabulary
	orm := detDb.Model(&model.Vocabulary{})

	// todo category
	if *option.Level != "" {
		orm = orm.Where("level = ?", option.Level)
	}

	if err := orm.Order("RANDOM()").Limit(30).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
