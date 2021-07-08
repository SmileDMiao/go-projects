package model

import "time"

// Movie struct
type Movie struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Other     string    `json:"other"`
	Desc      string    `json:"desc"`
	Year      string    `json:"year"`
	Area      string    `json:"area"`
	Tag       string    `json:"tag"`
	Star      string    `json:"star"`
	Comment   string    `json:"comment"`
	Quote     string    `json:"quote"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
