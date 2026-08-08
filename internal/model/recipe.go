package model

import "time"

type Recipe struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	MakingTime  string    `json:"making_time"`
	Serves      string    `json:"serves"`
	Ingredients string    `json:"ingredients"`
	Cost        int       `json:"cost"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
