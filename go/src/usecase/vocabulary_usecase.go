package vocabulary

import (
	"nothing-behind.com/sample_gin/db"
	"nothing-behind.com/sample_gin/entity"
)

type Usecase struct{}

type Vocabulary entity.Vocabulary

func (s Usecase) GetAll() ([]Vocabulary, error) {
	detDb := db.GetDB()
	var u []Vocabulary

	if err := detDb.Order("RANDOM()").Limit(30).Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
