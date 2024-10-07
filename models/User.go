package models

type User struct {
	ID             int    `gorm:"primary_key;column:id;type:int;not null"`
	Email          string `gorm:"unique;column:email;type:varchar(255);not null"`
	Name           string `gorm:"column:name;type:varchar(100);not null"`
	ProfilePicture []byte `gorm:"column:profile_picture;type:longblob"`
	Password       []byte `gorm:"column:password;type:varbinary(255);not null"`
	RoleID         int    `gorm:"index;column:role_id;type:int;not null"`
	Role		   Role   `gorm:"foreignKey:RoleID;references:ID"`
}

type UserResponse struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	ProfilePicture []byte `json:"profile_picture"`
	RoleID         int    `json:"role_id"`
}
