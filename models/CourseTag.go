package models

type CourseTag struct {
	IDCourse int `gorm:"primary_key;column:id_course;type:int;not null"`
	IDTag    int `gorm:"primary_key;column:id_tag;type:int;not null"`
	Votes    int `gorm:"column:votes;type:int"`
}

type CourseTagResponse struct {
	IDCourse int `json:"id_course"`
	IDTag    int `json:"id_tag"`
	Votes    int `json:"votes"`
}