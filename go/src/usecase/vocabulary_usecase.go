package vocabulary

import (
	"nothing-behind.com/sample_gin/db"
	"nothing-behind.com/sample_gin/entity"
)

// todo category
type ListInput struct {
	Level string
}

type Usecase struct{}

type Vocabulary entity.Vocabulary

func (s Usecase) GetAll(input *ListInput) ([]Vocabulary, error) {
	detDb := db.GetDB()
	var u []Vocabulary

	// todo category
	if err := detDb.Where("level = ?", input.Level).Order("RANDOM()").Limit(30).Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
