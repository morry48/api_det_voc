package repository

import (
	"nothing-behind.com/sample_gin/packages/vocabulary/domain/entity"
)

type VocabularyRepository interface {
	SelectByOption(option *entity.QueryOption) (*[]entity.Vocabulary, error)
}
