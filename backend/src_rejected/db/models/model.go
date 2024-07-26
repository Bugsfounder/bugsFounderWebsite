package models

import "time"

type Blog struct {
	Title     string    `gorm:"primaryKey" json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	Hidden    bool      `json:"is_hidden"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SubTutorial struct {
	Title     string    `gorm:"primaryKey" json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	Hidden    bool      `json:"is_hidden"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tutorial struct {
	SubTutorial *[]SubTutorial `gorm:"primaryKey" json:"sub_tutorial"`
	Title       string         `gorm:"primaryKey" json:"title"`
	Description string         `json:"description"`
	Author      string         `json:"author"`
	Tags        []string       `json:"tags"`
	Hidden      bool           `json:"is_hidden"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type User struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"primaryKey" json:"username"`
	Email     string `gorm:"primaryKey" json:"email"`
	Password  string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
