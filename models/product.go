package models

import (
	"time"
)

type Product struct {
	Id          int64     `json:"id"`
	Brand       string    `json:"brand"`
	Name        string    `json:"name"`
	CategoryId  int64     `json:"category_id"`
	Price       int64     `json:"price"`
	Description *string   `json:"description"`
	Sizes       []Size    `json:"sizes"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	Images      []Image   `json:"images"`
}

type Size struct {
	Size  string `json:"size"`
	Count int64  `json:"count"`
}
