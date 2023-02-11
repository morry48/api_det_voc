package usecase

import (
	"nothing-behind.com/sample_gin/packages/vocabulary/domain/entity"
	repository "nothing-behind.com/sample_gin/packages/vocabulary/domain/interface_repository"
	"nothing-behind.com/sample_gin/packages/vocabulary/infra/postgres"
)

type ListInput struct {
	Level *string
}

type ListOutput struct {
	Vocabularies *[]entity.Vocabulary
}

type ListVocabularies interface {
	Exec(input *ListInput) (*ListOutput, error)
}

type listVocabularies struct {
	db                   *postgres.Database
	vocabularyRepository repository.VocabularyRepository
}

func NewListCategories(db *postgres.Database, vocabularyRepository repository.VocabularyRepository) ListVocabularies {
	return &listVocabularies{
		db:                   db,
		vocabularyRepository: vocabularyRepository,
	}
}

func (s listVocabularies) Exec(params *ListInput) (*ListOutput, error) {
	option := entity.QueryOption{
		Level: params.Level,
	}

	vocabularyList, err := s.vocabularyRepository.SelectByOption(&option)
	if err != nil {
		return nil, err
	}

	listOutput := ListOutput{
		Vocabularies: vocabularyList,
	}
	return &listOutput, err

}
