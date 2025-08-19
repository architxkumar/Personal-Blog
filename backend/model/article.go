package model

import "github.com/google/uuid"

type Article struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    string    `json:"date"`
}

type ArticlePreviewDTO struct {
	Blogs []struct {
		Id    uuid.UUID `json:"id"`
		Title string    `json:"title"`
	} `json:"blogs"`
}
