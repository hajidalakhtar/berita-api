package model

import "time"

type NewsApiOrgResponse struct {
	Status       string       `json:"status"`
	TotalResults int          `json:"totalResults"`
	Articles     []NewsApiOrg `json:"articles"`
}

type NewsApiOrg struct {
	Source      Source    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	UrlToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type Source struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
