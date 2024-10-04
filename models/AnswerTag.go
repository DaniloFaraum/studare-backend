package models

type AnswerTag struct {
	IDTag    int `gorm:"primary_key;column:id_tag;type:int;not null"`
	IDAnswer int `gorm:"primary_key;column:id_answer;type:int;not null"`
}

type AnswerTagResponse struct {
	IDTag    int `json:"id_tag"`
	IDAnswer int `json:"id_answer"`
}
