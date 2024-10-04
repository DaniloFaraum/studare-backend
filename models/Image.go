package models

type Image struct {
	NameImage      string `gorm:"column:name_image;type:varchar(64)"`
	ID             int    `gorm:"primary_key;column:id;type:int;not null"`
	EncryptedImage []byte `gorm:"column:encrypted_image;type:longblob;not null"`
}

type ImageResponse struct {
	NameImage      string `json:"name_image"`
	ID             int    `json:"id"`
	EncryptedImage []byte `json:"encrypted_image"`
}