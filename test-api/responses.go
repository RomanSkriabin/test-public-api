package testapi

import (
	"errors"
	"fmt"
)

// Структуры
type Post struct {
	NumLikes          int    `json:"num_likes"`
	PostType          string `json:"post_type"`
	PostUrl           string `json:"post_url"`
	Posted            string `json:"posted"`
	PosterLinkedinUrl string `json:"poster_linkedin_url"`
	PosterName        string `json:"poster_name"`
	PosterTitle       string `json:"poster_title"`
	Text              string `json:"text"`
	Urn               string `json:"urn"`
}

type PostList struct {
	Items   []Post `json:"data"`
	Message string `json:"message"`
	Total   int    `json:"total"`
}

func (p Post) Info() string {
	return fmt.Sprintf("Title: %s ; Date: %s ; Likes: %d", p.PosterTitle, p.Posted, p.NumLikes)
}

type Recomendation struct {
	Profile_url string `json:"profile_url"`
	Text        string `json:"text"`
}

type RecomendationList struct {
	Items   []Recomendation `json:"data"`
	Message string          `json:"message"`
}

func (p RecomendationList) Info() (string, error) {
	result := ""
	if len(p.Items) == 0 {
		return "", errors.New("empty result")
	}
	if p.Message != "ok" {
		return "", errors.New("no ok request")
	}
	for _, item := range p.Items {
		result += fmt.Sprintf("Recomendation From: %s\n Text: %s\n----------------------\n", item.Profile_url, item.Text)
	}
	return result, nil
}
