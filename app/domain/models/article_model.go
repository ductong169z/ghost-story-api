package models

import (
	"database/sql"
	"gfly/app/domain/models/types"
	"time"

	mb "github.com/gflydev/db"
)

// ====================================================================
// ============================ Data Types ============================
// ====================================================================

// TBD

// ====================================================================
// ============================== Table ===============================
// ====================================================================

// TableArticle Table name
const TableArticle = "articles"

// Article struct to describe an article object.
type Article struct {
	// Table meta data
	MetaData mb.MetaData `db:"-" model:"table:articles"`

	// Table fields
	ID             int                 `db:"id" model:"name:id; type:serial,primary"`
	Title          string              `db:"title" model:"name:title"`
	Slug           string              `db:"slug" model:"name:slug"`
	Excerpt        sql.NullString      `db:"excerpt" model:"name:excerpt"`
	Content        string              `db:"content" model:"name:content"`
	CoverImage     sql.NullString      `db:"cover_image" model:"name:cover_image"`
	Status         types.ArticleStatus `db:"status" model:"name:status"`
	SEODescription sql.NullString      `db:"seo_description" model:"name:seo_description"`
	SEOKeywords    sql.NullString      `db:"seo_keywords" model:"name:seo_keywords"`
	AuthorID       int                 `db:"author_id" model:"name:author_id"`
	PublishedAt    sql.NullTime        `db:"published_at" model:"name:published_at"`
	YouTubeURL     sql.NullString      `db:"youtube_url" model:"name:youtube_url"`
	TikTokURL      sql.NullString      `db:"tiktok_url" model:"name:tiktok_url"`
	ViewCount      int                 `db:"view_count" model:"name:view_count"`
	CreatedAt      time.Time           `db:"created_at" model:"name:created_at"`
	UpdatedAt      sql.NullTime        `db:"updated_at" model:"name:updated_at"`
	DeletedAt      sql.NullTime        `db:"deleted_at" model:"name:deleted_at"`
}
