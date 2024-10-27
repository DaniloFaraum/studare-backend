package models

import (
	"time"
)

type Course struct {
	ID          int       `gorm:"primary_key;column:id;type:int;not null"`
	Name        string    `gorm:"column:name;type:varchar(255);not null"`
	Description string    `gorm:"column:description;type:text"`
	Link        string    `gorm:"column:link;type:varchar(255)"`
	Rating      float64       `gorm:"column:rating;type:float"`
	Duration    time.Time `gorm:"column:duration;type:time"`
	Author      string    `gorm:"column:author;type:varchar(100)"`
	Institution string    `gorm:"column:institution;type:varchar(100)"`
	IDImage     int       `gorm:"index;column:id_image;type:int"`
	Image       Image     `gorm:"foreignKey:IDImage;references:ID"`
	Tags        []Tag     `gorm:"many2many:course_tags"`
}

type CourseResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Rating      float64   `json:"rating"`
	Duration    time.Time `json:"duration"`
	Author      string    `json:"author"`
	Institution string    `json:"institution"`
	IDImage     int       `json:"id_image"`
}
