package vocabulary

import (
	"nothing-behind.com/sample_gin/db"
	"nothing-behind.com/sample_gin/entity"
)

type Usecase struct{}

type Vocabulary entity.Vocabulary

func (s Usecase) GetAll() ([]Vocabulary, error) {
	db := db.GetDB()
	var u []Vocabulary

	if err := db.Limit(30).Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
