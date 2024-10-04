package models

type QuestionTag struct {
	IDTag      int `gorm:"primary_key;column:id_tag;type:int;not null"`
	IDQuestion int `gorm:"primary_key;column:id_question;type:int;not null"`
}

type QuestionTagResponse struct {
	IDTag      int `json:"id_tag"`
	IDQuestion int `json:"id_question"`
}