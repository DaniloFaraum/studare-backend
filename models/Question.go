package models

type Question struct {
	ID              int           `gorm:"primary_key;column:id;type:int;not null"`
	IDQuestionnaire int           `gorm:"index;column:id_questionnaire;type:int;not null"`
	Question        string        `gorm:"column:question;type:text;not null"`
	Questionnaire   Questionnaire `gorm:"foreignKey:IDQuestionnaire;references:ID"`
	Tags            []Tag         `gorm:"many2many:question_tags"`
}

type QuestionResponse struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
}
