package models

type Questionnaire struct {
    ID     int    `gorm:"primary_key;column:id;type:int;not null"`
    IDUser int    `gorm:"index;column:id_user;type:int;not null"`
    Name   string `gorm:"column:name;type:varchar(255);not null"`
    Ready  int8   `gorm:"column:ready;type:tinyint(1);not null"`
    User   User   `gorm:"foreignKey:IDUser;references:ID"`
}

type QuestionnaireResponse struct {
    ID     int    `json:"id"`
    IDUser int    `json:"id_user"`
    Name   string `json:"name"`
    Ready  int8   `json:"ready"`
}