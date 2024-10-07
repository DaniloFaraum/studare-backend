package models

type Rating struct {
	ID         int    `gorm:"primary_key;column:id;type:int;not null"`
	IDUser     int    `gorm:"index;column:id_user;type:int;not null"`
	IDCourse   int    `gorm:"index;column:id_course;type:int;not null"`
	Opinion    int    `gorm:"column:opinion;type:int;not null"`
	Commentary string `gorm:"column:commentary;type:varchar(500)"`
	User       User   `gorm:"foreignKey:IDUser;references:ID"`
	Course     Course `gorm:"foreignKey:IDCourse;references:ID"`
	Tags       []Tag  `gorm:"many2many:rating_tags"`
}

type RatingResponse struct {
	ID         int    `json:"id"`
	IDUser     int    `json:"id_user"`
	IDCourse   int    `json:"id_course"`
	Opinion    int    `json:"opinion"`
	Commentary string `json:"commentary"`
}
