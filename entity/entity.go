package entity

import "time"

type Users struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Images struct {
	Id         int       `json:"id" db:"id"`
	UserId     int       `json:"user_id" db:"user_id"`
	ImageURL   string    `json:"image_url" db:"image_url"`
	UploadTime time.Time `json:"upload_time" db:"upload_time"`
}
