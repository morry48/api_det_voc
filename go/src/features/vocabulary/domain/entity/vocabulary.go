package entity

type Vocabulary struct {
	ID       uint   `json:"id"`
	EnWord   string `json:"en_word"`
	JpWord   string `json:"jp_word"`
	Category string `json:"category"`
	Level    string `json:"level"`
}

type QueryOption struct {
	Level *string
}
