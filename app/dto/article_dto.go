package dto

import "gfly/app/domain/models/types"

// CreateArticle struct to describe the request body to create a new article.
// @Description Request payload for creating a new article.
// @Tags Articles
type CreateArticle struct {
	Title          string              `json:"title" example:"How to Build a Go Web Application" validate:"required,max=255" doc:"Article title (required, max length 255)"`
	Slug           string              `json:"slug" example:"how-to-build-go-web-application" validate:"required,max=255" doc:"URL-friendly slug (required, max length 255)"`
	Excerpt        string              `json:"excerpt" example:"Learn how to build a web application using Go" validate:"omitempty" doc:"Short excerpt/summary of the article (optional)"`
	Content        string              `json:"content" example:"<p>This is the full content of the article...</p>" validate:"required" doc:"Full HTML content of the article (required)"`
	CoverImage     string              `json:"cover_image" example:"https://example.com/images/cover.jpg" validate:"omitempty,max=255" doc:"URL of the article cover image (optional, max length 255)"`
	Status         types.ArticleStatus `json:"status" example:"draft" validate:"omitempty,oneof=draft published archived" doc:"Article status (optional, one of: draft, published, archived)"`
	SEODescription string              `json:"seo_description" example:"Comprehensive guide to building Go web applications" validate:"omitempty,max=300" doc:"SEO meta description (optional, max length 300)"`
	SEOKeywords    string              `json:"seo_keywords" example:"golang,web development,tutorial" validate:"omitempty" doc:"SEO keywords (optional)"`
	AuthorID       int                 `json:"author_id" example:"1" validate:"required" doc:"ID of the article author (required)"`
	YouTubeURL     string              `json:"youtube_url" example:"https://youtube.com/watch?v=abcdef" validate:"omitempty,max=255" doc:"YouTube video URL (optional, max length 255)"`
	TikTokURL      string              `json:"tiktok_url" example:"https://tiktok.com/@user/video/123456" validate:"omitempty,max=255" doc:"TikTok video URL (optional, max length 255)"`
}

// UpdateArticle struct to partially update an existing article.
// @Description Request payload for updating an existing article.
// @Tags Articles
type UpdateArticle struct {
	ID             int                 `json:"-" validate:"omitempty,gte=1" doc:"Article ID (greater than or equal to 1)"`
	Title          string              `json:"title" example:"Updated: How to Build a Go Web Application" validate:"omitempty,max=255" doc:"Updated article title (optional, max length 255)"`
	Slug           string              `json:"slug" example:"updated-how-to-build-go-web-application" validate:"omitempty,max=255" doc:"Updated URL-friendly slug (optional, max length 255)"`
	Excerpt        string              `json:"excerpt" example:"Updated summary of building a Go web application" validate:"omitempty" doc:"Updated excerpt/summary (optional)"`
	Content        string              `json:"content" example:"<p>Updated content of the article...</p>" validate:"omitempty" doc:"Updated HTML content (optional)"`
	CoverImage     string              `json:"cover_image" example:"https://example.com/images/updated-cover.jpg" validate:"omitempty,max=255" doc:"Updated cover image URL (optional, max length 255)"`
	Status         types.ArticleStatus `json:"status" example:"published" validate:"omitempty,oneof=draft published archived" doc:"Updated article status (optional, one of: draft, published, archived)"`
	SEODescription string              `json:"seo_description" example:"Updated guide to building Go web applications" validate:"omitempty,max=300" doc:"Updated SEO description (optional, max length 300)"`
	SEOKeywords    string              `json:"seo_keywords" example:"updated,golang,web development" validate:"omitempty" doc:"Updated SEO keywords (optional)"`
	YouTubeURL     string              `json:"youtube_url" example:"https://youtube.com/watch?v=updated" validate:"omitempty,max=255" doc:"Updated YouTube URL (optional, max length 255)"`
	TikTokURL      string              `json:"tiktok_url" example:"https://tiktok.com/@user/video/updated" validate:"omitempty,max=255" doc:"Updated TikTok URL (optional, max length 255)"`
}

// UpdateArticleStatus struct allows update `status` field from an existing article.
// @Description Request payload for updating the status field of an existing article.
// @Tags Articles
type UpdateArticleStatus struct {
	ID     int                 `json:"-" validate:"omitempty" doc:"Article ID associated with the status update"`
	Status types.ArticleStatus `json:"status" example:"published" validate:"required,oneof=draft published archived" doc:"New status of the article (required, one of: draft, published, archived)"`
}

type ArticleFilter struct {
	Filter
	Status types.ArticleStatus `json:"status" example:"published" validate:"omitempty,oneof=draft published archived" doc:"Article status (optional, one of: draft, published, archived)"`
}
