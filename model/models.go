package models

import (
	"time"
)

type AboutMe struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Skills      []string  `json:"skills"`
	IsSelected  bool      `json:"is_selected"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Certificate struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Authority  string    `json:"authority"`
	Credential string    `json:"credential"`
	URL        string    `json:"url"`
	Image      string    `json:"image"`
	DateEarned time.Time `json:"date_earned"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type SocialMedia struct {
	ID        string    `json:"id"`
	Platform  string    `json:"platform"`
	Username  string    `json:"username"`
	URL       string    `json:"url"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Project struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Technologies []string  `json:"technologies"`
	URL          string    `json:"url"`
	Image        string    `json:"image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Blog struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Journey struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Technology struct {
	ID        string    `json:"id"`
	Icon      string    `json:"icon"`
	Name      string    `json:"name"`
	Skill     int       `json:"skill"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
