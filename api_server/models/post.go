package models

import (
	"encoding/json"
)

type Post struct {
	Title      string
	Slug       string
	URL        string
	Content    string
	Date       string
	Image      string
	CategoryID Category
	TotalPosts int
}

type PostJSON struct {
	Title      string   `json:"title, omitempty"`
	Slug       string   `json:"slug, omitempty"`
	URL        string   `json:"url, omitempty"`
	Content    string   `json:"content"`
	Date       string   `json:"date, omitempty"`
	Image      string   `json:"image"`
	CategoryID Category `json:"category_id, omitempty"`
	TotalPosts int      `json:"total_posts"`
}

func (p *Post) MarshalJSON() ([]byte, error) {
	return json.Marshal(PostJSON{
		p.Title,
		p.Slug,
		p.URL,
		p.Content,
		p.Date,
		p.Image,
		p.CategoryID,
		p.TotalPosts,
	})
}

func (p *Post) UnmarshalJSON(b []byte) error {
	temp := &PostJSON{}

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	p.Title = temp.Title
	p.Slug = temp.Slug
	p.URL = temp.URL
	p.Content = temp.Content
	p.Date = temp.Date
	p.Image = temp.Image
	p.CategoryID = temp.CategoryID
	p.TotalPosts = temp.TotalPosts

	return nil
}
