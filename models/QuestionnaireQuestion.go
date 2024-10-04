package models

type QuestionnaireQuestion struct {
	IDQuestionnaire int `gorm:"primary_key;column:id_questionnaire;type:int;not null"`
	IDQuestion      int `gorm:"primary_key;column:id_question;type:int;not null"`
}

type QuestionnaireQuestionResponse struct {
	IDQuestionnaire int `json:"id_questionnaire"`
	IDQuestion      int `json:"id_question"`
}