package models

type Comment struct {
	ID       int    `gorm:"primary_key;column:id;type:int;not null"`
	IDCourse int    `gorm:"index;column:id_course;type:int;not null"`
	IDUser   int    `gorm:"index;column:id_user;type:int;not null"`
	Content  string `gorm:"column:content;type:varchar(500)"`
	Likes    int    `gorm:"column:likes;type:int"`
	Dislikes int    `gorm:"column:dislikes;type:int"`
	Course   Course `gorm:"foreignKey:IDCourse;references:ID"`
	User     User   `gorm:"foreignKey:IDUser;references:ID"`
}

type CommentResponse struct {
	ID       int    `json:"id"`
	IDCourse int    `json:"id_course"`
	IDUser   int    `json:"id_user"`
	Content  string `json:"content"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}
