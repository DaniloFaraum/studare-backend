package models

type QuestionTag struct {
    TagID      int `gorm:"primaryKey;column:tag_id;not null"`
    QuestionID int `gorm:"primaryKey;column:question_id;not null"`
}

type QuestionTagResponse struct {
    TagID      int `json:"tag_id"`
    QuestionID int `json:"question_id"`
}