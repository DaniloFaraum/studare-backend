package models

type Role struct {
	ID   int    `gorm:"primary_key;column:id;type:int;not null"`
	Name string `gorm:"column:name;type:varchar(50);not null"`
}

type RoleResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
