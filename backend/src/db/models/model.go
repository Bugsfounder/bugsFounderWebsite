package models

import "time"

type Blog struct {
	BlogID    int       `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"primaryKey" json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	Hidden    bool      `json:"is_hidden"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tutorial struct {
	TutorialID int       `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"primaryKey" json:"title"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
	Tags       []string  `json:"tags"`
	Hidden     bool      `json:"is_hidden"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Tutorials struct {
	TutorialID  *[]Tutorial `gorm:"primaryKey" json:"tutorial_id"`
	Title       string      `gorm:"primaryKey" json:"title"`
	Description string      `json:"description"`
	Author      string      `json:"author"`
	Tags        []string    `json:"tags"`
	Hidden      bool        `json:"is_hidden"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type User struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"primaryKey" json:"username"`
	Email     string `gorm:"primaryKey" json:"email"`
	Password  string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
