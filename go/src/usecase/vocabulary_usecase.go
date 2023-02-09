package vocabulary

import (
	"nothing-behind.com/sample_gin/db"
	"nothing-behind.com/sample_gin/entity"
)

// todo category
type ListInput struct {
	Level *string
}

type Usecase struct{}

type Vocabulary entity.Vocabulary

// todo category
func (s Usecase) GetAll(input *ListInput) ([]Vocabulary, error) {
	detDb := db.GetDB()
	var u []Vocabulary
	orm := detDb.Model(&Vocabulary{})

	if *input.Level != "" {
		orm = orm.Where("level = ?", input.Level)
	}

	if err := orm.Order("RANDOM()").Limit(30).Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
