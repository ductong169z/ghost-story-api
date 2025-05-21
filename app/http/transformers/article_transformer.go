package transformers

import (
	"gfly/app/domain/models"
	"gfly/app/http/response"
)

// ToArticleResponse transforms an Article model to an Article response
func ToArticleResponse(article models.Article) response.Article {
	return response.Article{
		ID:             article.ID,
		Title:          article.Title,
		Slug:           article.Slug,
		Excerpt:        article.Excerpt.String,
		Content:        article.Content,
		CoverImage:     article.CoverImage.String,
		Status:         string(article.Status),
		SEODescription: article.SEODescription.String,
		SEOKeywords:    article.SEOKeywords.String,
		AuthorID:       article.AuthorID,
		PublishedAt:    article.PublishedAt.Time,
		YouTubeURL:     article.YouTubeURL.String,
		TikTokURL:      article.TikTokURL.String,
		ViewCount:      article.ViewCount,
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt.Time,
	}
}

// ToArticleListResponse transforms a slice of Article models to a slice of Article responses
func ToArticleListResponse(articles []models.Article) []response.Article {
	result := make([]response.Article, len(articles))
	for i, article := range articles {
		result[i] = ToArticleResponse(article)
	}
	return result
}

// ToArticleResponse transforms an Article model to an Article response
func ToArticleForGuestResponse(article models.Article) response.Article {
	return response.Article{
		Title:          article.Title,
		Slug:           article.Slug,
		Excerpt:        article.Excerpt.String,
		Content:        article.Content,
		CoverImage:     article.CoverImage.String,
		Status:         string(article.Status),
		SEODescription: article.SEODescription.String,
		SEOKeywords:    article.SEOKeywords.String,
		AuthorID:       article.AuthorID,
		PublishedAt:    article.PublishedAt.Time,
		YouTubeURL:     article.YouTubeURL.String,
		TikTokURL:      article.TikTokURL.String,
		ViewCount:      article.ViewCount,
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt.Time,
	}
}

// ToArticleListResponse transforms a slice of Article models to a slice of Article responses
func ToArticleListForGuestResponse(articles []models.Article) []response.Article {
	result := make([]response.Article, len(articles))
	for i, article := range articles {
		result[i] = ToArticleForGuestResponse(article)
	}
	return result
}
