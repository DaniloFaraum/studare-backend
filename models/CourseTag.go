package models

type CourseTag struct {
    TagID    int `gorm:"primaryKey;column:tag_id;not null"`
    CourseID int `gorm:"primaryKey;column:course_id;not null"`
}

type CourseTagResponse struct {
    TagID    int `json:"tag_id"`
    CourseID int `json:"course_id"`
    Votes    int `json:"votes"`
}