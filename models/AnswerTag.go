package models

type AnswerTag struct {
    TagID    int `gorm:"primaryKey;column:tag_id;not null"`
    AnswerID int `gorm:"primaryKey;column:answer_id;not null"`
}

type AnswerTagResponse struct {
    TagID    int `json:"tag_id"`
    AnswerID int `json:"answer_id"`
}