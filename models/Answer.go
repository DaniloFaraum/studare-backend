package models

type Answer struct {
	ID         int      `gorm:"primary_key;column:id;type:int;not null"`
	IDQuestion int      `gorm:"index;column:id_question;type:int;not null"`
	Text       string   `gorm:"column:text;type:varchar(255);not null"`
	Voted      int8     `gorm:"column:voted;type:tinyint(1)"`
	Question   Question `gorm:"foreignKey:IDQuestion;references:ID"`
	Tags       []Tag    `gorm:"many2many:answer_tags"`
}

type AnswerResponse struct {
	ID         int    `json:"id"`
	IDQuestion int    `json:"id_question"`
	Text       string `json:"text"`
	Voted      int8   `json:"voted"`
}
