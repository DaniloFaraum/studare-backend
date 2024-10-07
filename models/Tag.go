package models

type Tag struct {
    ID        int       `gorm:"primary_key;column:id;type:int;not null"`
    Name      string    `gorm:"column:name;type:varchar(255);not null"`
    Courses   []Course  `gorm:"many2many:course_tags"`
    Questions []Question `gorm:"many2many:question_tags"`
    Answers   []Answer  `gorm:"many2many:answer_tags"`
    Ratings   []Rating  `gorm:"many2many:rating_tags"`
}

type TagResponse struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}