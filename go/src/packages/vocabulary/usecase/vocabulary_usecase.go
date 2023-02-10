package vocabulary

import (
	"github.com/jinzhu/gorm"
	"nothing-behind.com/sample_gin/packages/vocabulary/domain/entity"
	repository2 "nothing-behind.com/sample_gin/packages/vocabulary/domain/interface_repository"
	"nothing-behind.com/sample_gin/packages/vocabulary/infra/postgres/model"
	"nothing-behind.com/sample_gin/packages/vocabulary/infra/postgres/repository"
)

type ListInput struct {
	Level *string
}

type ListOutput struct {
	Vocabularies []Vocabulary
}

type Usecase struct{}

type Vocabulary model.Vocabulary

func (s Usecase) GetListByParams(params *ListInput) (*[]entity.Vocabulary, error) {
	option := entity.QueryOption{
		Level: params.Level,
	}

	// todo 依存解決をライブラリにする(現状だとユースケースで依存解決を行っているためインフラの知識が漏れ出ている)
	var db *gorm.DB
	var vocRepo repository2.VocabularyRepository
	vocRepo = repository.NewVocabularyRepository(db)

	output, err := vocRepo.SelectByOption(&option)

	if err != nil {
		return nil, err
	}
	return output, err

}
