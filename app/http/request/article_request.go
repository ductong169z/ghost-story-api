package request

import "gfly/app/dto"

// ====================================================================
// ========================== Add Requests ============================
// ====================================================================

// ---------------------- Create Article ------------------------

type CreateArticle struct {
	dto.CreateArticle
}

// ToDto Convert to CreateArticle DTO object.
func (r CreateArticle) ToDto() dto.CreateArticle {
	return r.CreateArticle
}

// ====================================================================
// ========================= Update Requests ==========================
// ====================================================================

// ---------------------- Update Article ------------------------

type UpdateArticle struct {
	dto.UpdateArticle
}

// ToDto Convert to UpdateArticle DTO object.
func (r UpdateArticle) ToDto() dto.UpdateArticle {
	return r.UpdateArticle
}

func (r UpdateArticle) SetID(id int) {
	r.ID = id
}

// ------------------ Update Article's status --------------------

// UpdateArticleStatus struct to describe update article's status
type UpdateArticleStatus struct {
	dto.UpdateArticleStatus
}

func (r UpdateArticleStatus) SetID(id int) {
	r.ID = id
}

// ToDto convert struct to UpdateArticleStatus DTO object
func (r UpdateArticleStatus) ToDto() dto.UpdateArticleStatus {
	return r.UpdateArticleStatus
}
