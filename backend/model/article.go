package model

import "github.com/google/uuid"

type Article struct {
	Id      uuid.UUID `json:"id"`
	Content string    `json:"content"`
	Date    string    `json:"date"`
}
