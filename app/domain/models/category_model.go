package models

import (
	"database/sql"
	"time"

	mb "github.com/gflydev/db"
)

// TableCategory Table name
const TableCategory = "categories"

// Category struct to describe a category object.
type Category struct {
	// Table meta data
	MetaData mb.MetaData `db:"-" model:"table:categories"`

	// Table fields
	ID          int            `db:"id" model:"name:id; type:serial,primary"`
	Name        string         `db:"name" model:"name:name"`
	Slug        string         `db:"slug" model:"name:slug"`
	Description sql.NullString `db:"description" model:"name:description"`
	ParentID    sql.NullInt64  `db:"parent_id" model:"name:parent_id"`
	IsActive    bool           `db:"is_active" model:"name:is_active"`
	SortOrder   int            `db:"sort_order" model:"name:sort_order"`
	CreatedAt   time.Time      `db:"created_at" model:"name:created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at" model:"name:updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at" model:"name:deleted_at"`

	// Relations
	Parent *Category `db:"-"`
}

// TableName returns the table name for the Category model
func (Category) TableName() string {
	return TableCategory
}

// BeforeCreate handles pre-creation logic
func (c *Category) BeforeCreate() error {
	c.CreatedAt = time.Now()
	return nil
}

// BeforeUpdate handles pre-update logic
func (c *Category) BeforeUpdate() error {
	now := time.Now()
	c.UpdatedAt = sql.NullTime{Time: now, Valid: true}
	return nil
}
