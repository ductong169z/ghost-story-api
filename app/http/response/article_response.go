package response

import "time"

// Article response structure for API
type Article struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Slug           string    `json:"slug"`
	Excerpt        string    `json:"excerpt,omitempty"`
	Content        string    `json:"content"`
	CoverImage     string    `json:"cover_image,omitempty"`
	Status         string    `json:"status"`
	SEODescription string    `json:"seo_description,omitempty"`
	SEOKeywords    string    `json:"seo_keywords,omitempty"`
	AuthorID       int       `json:"author_id"`
	PublishedAt    time.Time `json:"published_at,omitempty"`
	YouTubeURL     string    `json:"youtube_url,omitempty"`
	TikTokURL      string    `json:"tiktok_url,omitempty"`
	ViewCount      int       `json:"view_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
