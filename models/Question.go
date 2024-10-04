package models

type Question struct {
	ID       int    `gorm:"primary_key;column:id;type:int;not null"`
	Question string `gorm:"column:question;type:text;not null"`
}

type QuestionResponse struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
}