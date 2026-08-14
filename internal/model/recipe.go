package model

import "time"

type Recipe struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	MakingTime  string    `json:"making_time" binding:"required"`
	Serves      string    `json:"serves" binding:"required"`
	Ingredients string    `json:"ingredients" binding:"required"`
	Cost        int       `json:"cost" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
