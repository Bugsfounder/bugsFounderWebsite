package models

import (
	"time"

	"github.com/bugsfounder/bugsfounderweb/admin"
)

type Blog struct {
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Author    string          `json:"author"`
	Tags      []string        `json:"tags"`
	CreatedAt time.Time       `json:"created_at"`
	Url       string          `json:"url"`
	IsAdmin   bool            `json:"is_admin"`
	AdminMode admin.AdminMode `json:"admin_mode"`
	IsHidden  bool            `json:"is_hidden"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type Tutorial struct {
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	Author        string          `json:"author"`
	Tags          []string        `json:"tags"`
	Sub_Tutorials []Sub_Tutorial  `json:"sub_tutorials"`
	Url           string          `json:"url"`
	IsAdmin       bool            `json:"is_admin"`
	AdminMode     admin.AdminMode `json:"admin_mode"`
	IsHidden      bool            `json:"is_hidden"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type Sub_Tutorial struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BlacklistedToken struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Admin struct {
	Email     string          `json:"email"`
	Username  string          `json:"username"`
	Password  string          `json:"password"`
	AdminMode admin.AdminMode `json:"adminMode"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type UPDATE_USER_INFO int

const (
	UPDATE_PASSWORD UPDATE_USER_INFO = iota + 1
	UPDATE_EMAIL
	UPDATE_USERNAME
)

type UpdateUser struct {
	Email     string           `json:"email"`
	Username  string           `json:"username"`
	Password  string           `json:"password"`
	Update    UPDATE_USER_INFO `json:"update"`
	UpdatedAt time.Time        `json:"updated_at"`
}
