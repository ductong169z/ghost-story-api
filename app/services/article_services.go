package services

import (
	"fmt"
	"gfly/app/domain/models"
	"gfly/app/domain/models/types"
	"gfly/app/dto"
	"slices"
	"strings"
	"time"

	"github.com/gflydev/core"
	"github.com/gflydev/core/errors"
	"github.com/gflydev/core/log"
	mb "github.com/gflydev/db"
	dbNull "github.com/gflydev/db/null"
	qb "github.com/jivegroup/fluentsql"
)

// ====================================================================
// ========================= Main functions ===========================
// ====================================================================

// FindArticles retrieves a list of articles from the database based on the provided filter criteria.
// It supports searching by keyword, ordering by specified fields, and pagination.
//
// Parameters:
//   - filterDto (dto.Filter): The filter containing search criteria, order by field, page, and per-page details.
//
// Returns:
//
//	([]models.Article, int, error): A list of article models, the total number of articles, and any error encountered.
func FindArticles(filterDto dto.ArticleFilter) ([]models.Article, int, error) {
	// DB Model instance
	dbInstance := mb.Instance()
	// Error variable
	var err error

	// Define Article variable.
	var articles []models.Article
	var total int
	var offset = 0

	if filterDto.Page > 0 {
		offset = (filterDto.Page - 1) * filterDto.PerPage
	}

	builder := dbInstance.Select("*").
		Where(models.TableArticle+".deleted_at", qb.Null, nil).
		When(filterDto.Keyword != "", func(query qb.WhereBuilder) *qb.WhereBuilder {
			query.WhereGroup(func(queryGroup qb.WhereBuilder) *qb.WhereBuilder {
				queryGroup.Where(models.TableArticle+".title", qb.Like, "%"+filterDto.Keyword+"%").
					WhereOr(models.TableArticle+".slug", qb.Like, "%"+filterDto.Keyword+"%").
					WhereOr(models.TableArticle+".excerpt", qb.Like, "%"+filterDto.Keyword+"%").
					WhereOr(models.TableArticle+".content", qb.Like, "%"+filterDto.Keyword+"%")

				// Check if keyword matches any article status
				for _, status := range types.ArticleStatusList {
					if string(status) == filterDto.Keyword {
						queryGroup.WhereOr(models.TableArticle+".status", qb.Eq, filterDto.Status)
						break
					}
				}

				return &queryGroup
			})

			return &query
		}).
		Limit(filterDto.PerPage, offset)

	if filterDto.OrderBy != "" {
		// Default order by
		direction := qb.Asc
		orderKey := filterDto.OrderBy

		if strings.HasPrefix(filterDto.OrderBy, "-") {
			orderKey = filterDto.OrderBy[1:]
			direction = qb.Desc
		}

		var orderByFields = core.Data{
			"id":           fmt.Sprintf("%s.id", models.TableArticle),
			"title":        fmt.Sprintf("%s.title", models.TableArticle),
			"slug":         fmt.Sprintf("%s.slug", models.TableArticle),
			"status":       fmt.Sprintf("%s.status", models.TableArticle),
			"published_at": fmt.Sprintf("%s.published_at", models.TableArticle),
			"created_at":   fmt.Sprintf("%s.created_at", models.TableArticle),
		}

		if field, ok := orderByFields[orderKey]; ok {
			builder.OrderBy(field.(string), direction)
		}
	}

	// Query data
	total, err = builder.Find(&articles)

	// Return query result.
	return articles, total, err
}

// CreateArticle creates a new article in the system.
//
// This function performs the following steps:
// 1. Verifies that no other article exists with the same slug.
// 2. Creates a new article entity in the database.
//
// Parameters:
//   - createArticleDto (dto.CreateArticle): The payload containing the article details.
//
// Returns:
//   - (*models.Article, error): The created article object or an error if any step fails.
func CreateArticle(createArticleDto dto.CreateArticle) (*models.Article, error) {
	// Check if an article with the same slug already exists
	existingArticle, err := mb.GetModel[models.Article](qb.Condition{
		Field: models.TableArticle + ".slug",
		Opt:   qb.Eq,
		Value: createArticleDto.Slug,
	})

	if err == nil && existingArticle != nil {
		return nil, errors.New("An article with this slug already exists")
	}

	// Create new article
	article := &models.Article{
		Title:     createArticleDto.Title,
		Slug:      createArticleDto.Slug,
		Content:   createArticleDto.Content,
		AuthorID:  createArticleDto.AuthorID,
		Status:    createArticleDto.Status,
		CreatedAt: time.Now(),
	}

	// Set optional fields if provided
	if createArticleDto.Excerpt != "" {
		article.Excerpt = dbNull.String(createArticleDto.Excerpt)
	}

	if createArticleDto.CoverImage != "" {
		article.CoverImage = dbNull.String(createArticleDto.CoverImage)
	}

	if createArticleDto.SEODescription != "" {
		article.SEODescription = dbNull.String(createArticleDto.SEODescription)
	}

	if createArticleDto.SEOKeywords != "" {
		article.SEOKeywords = dbNull.String(createArticleDto.SEOKeywords)
	}

	if createArticleDto.YouTubeURL != "" {
		article.YouTubeURL = dbNull.String(createArticleDto.YouTubeURL)
	}

	if createArticleDto.TikTokURL != "" {
		article.TikTokURL = dbNull.String(createArticleDto.TikTokURL)
	}

	// Set published date if status is published
	if article.Status == types.ArticleStatusPublished {
		article.PublishedAt = dbNull.Time(time.Now())
	}

	// Create article in database
	if err := mb.CreateModel(article); err != nil {
		log.Errorf("Error while creating article: %v", err)
		return nil, errors.New("Error occurs while creating article")
	}

	return article, nil
}

// GetArticleByID retrieves an article by its ID.
//
// Parameters:
//   - articleID (int): The ID of the article to retrieve.
//
// Returns:
//   - (*models.Article, error): The article object or an error if not found.
func GetArticleByID(articleID int) (*models.Article, error) {
	article, err := mb.GetModelByID[models.Article](articleID)
	if err != nil {
		return nil, errors.New("Article not found")
	}

	return article, nil
}

// GetArticleBySlug retrieves an article by its slug and increments its view count.
//
// Parameters:
//   - slug (string): The slug of the article to retrieve.
//
// Returns:
//   - (*models.Article, error): The article object or an error if not found.
func GetArticleBySlug(slug string) (*models.Article, error) {
	article, err := mb.GetModel[models.Article](qb.Condition{
		Field: models.TableArticle + ".slug",
		Opt:   qb.Eq,
		Value: slug,
	})

	if err != nil {
		return nil, errors.New("Article not found")
	}

	// Increment view count when article is retrieved by slug
	// This is typically for guest/public views
	if err := IncrementArticleViewCount(article.ID); err != nil {
		log.Warnf("Failed to increment view count for article %d: %v", article.ID, err)
		// Continue even if view count update fails
	}

	return article, nil
}

// UpdateArticle updates an existing article in the system.
//
// This function fetches the article by its ID, updates the fields based on the given DTO.
//
// Parameters:
//   - updateArticleDto (dto.UpdateArticle): The DTO containing the article update data.
//
// Returns:
//   - (*models.Article, error): The updated article object or an error if any step fails.
//
// Possible Errors:
//   - "Article not found": Returned when no article is found for the provided ID.
//   - "Error occurs while updating article": Returned when an error occurs during the update process.
func UpdateArticle(updateArticleDto dto.UpdateArticle) (*models.Article, error) {
	// Get article by ID
	article, err := mb.GetModelByID[models.Article](updateArticleDto.ID)
	if err != nil {
		return nil, errors.New("Article not found")
	}

	// Check if slug is being updated and if it already exists
	if updateArticleDto.Slug != "" && updateArticleDto.Slug != article.Slug {
		existingArticle, err := mb.GetModel[models.Article](qb.Condition{
			Field: models.TableArticle + ".slug",
			Opt:   qb.Eq,
			Value: updateArticleDto.Slug,
		})

		if err == nil && existingArticle != nil && existingArticle.ID != article.ID {
			return nil, errors.New("An article with this slug already exists")
		}
	}

	// Update article with data from DTO
	article = updateArticleFromDto(article, updateArticleDto)

	// Update article in database
	if err := mb.UpdateModel(article); err != nil {
		log.Errorf("Error while updating article: %v", err)
		return nil, errors.New("Error occurs while updating article")
	}

	return article, nil
}

// UpdateArticleStatus updates the status of an existing article in the system.
//
// This function performs the following steps:
// 1. Finds the article by its ID.
// 2. Validates the new status value against allowed article states.
// 3. Updates the article status and related fields (like published_at if status is published).
//
// Parameters:
//   - updateArticleStatusDto (dto.UpdateArticleStatus): The DTO containing the article ID and new status.
//
// Returns:
//   - (*models.Article, error): The updated article object or an error if any step fails.
//
// Possible Errors:
//   - "Article not found": Returned when no article is found for the provided ID.
//   - "Error occurs while updating article's status": Returned when the provided status is invalid or the update process fails.
func UpdateArticleStatus(updateArticleStatusDto dto.UpdateArticleStatus) (*models.Article, error) {
	article, err := mb.GetModelByID[models.Article](updateArticleStatusDto.ID)
	if err != nil {
		return nil, errors.New("Article not found")
	}

	// Check article's status
	if !slices.Contains(types.ArticleStatusList, updateArticleStatusDto.Status) {
		return nil, errors.New("Error occurs while updating article's status: invalid status %v", updateArticleStatusDto.Status)
	}

	// Set new status
	article.Status = updateArticleStatusDto.Status

	// Update published_at if status is changed to published
	if article.Status == types.ArticleStatusPublished && !article.PublishedAt.Valid {
		article.PublishedAt = dbNull.Time(time.Now())
	}

	// Update article
	if err = mb.UpdateModel(article); err != nil {
		log.Errorf("Error while updating article status: %v", err)
		return nil, errors.New("Error occurs while updating article status")
	}

	return article, nil
}

// DeleteArticleByID deletes an article from the system.
//
// This function performs the following steps:
// 1. Fetches the article by its ID.
// 2. Deletes the article from the database.
//
// Parameters:
//   - articleID (int): The unique identifier of the article to be deleted.
//
// Returns:
//   - error: An error object if any step fails. Possible errors include:
//   - "Article not found": Returned when no article is found for the provided ID.
//   - "Error occurs while deleting article": Returned when an error occurs during the deletion process.
func DeleteArticleByID(articleID int) error {
	article, err := mb.GetModelByID[models.Article](articleID)
	if err != nil {
		return errors.New("Article not found")
	}

	// Delete article
	if err := mb.DeleteModel(article); err != nil {
		log.Errorf("Error while deleting article: %v", err)
		return errors.New("Error occurs while deleting article")
	}

	return nil
}

// IncrementArticleViewCount increments the view count for an article.
//
// Parameters:
//   - articleID (int): The ID of the article to update.
//
// Returns:
//   - error: An error object if the update fails.
func IncrementArticleViewCount(articleID int) error {
	// Get the current article
	article, err := mb.GetModelByID[models.Article](articleID)
	if err != nil {
		return errors.New("Article not found")
	}

	// Increment view count
	article.ViewCount++

	// Update only the view count field
	if err := mb.UpdateModel(article); err != nil {
		log.Errorf("Error while updating article view count: %v", err)
		return errors.New("Error occurs while updating article view count")
	}

	return nil
}

// ====================================================================
// ======================== Helper Functions ==========================
// ====================================================================

// updateArticleFromDto updates an existing Article model with data from UpdateArticle DTO.
// Only updates fields that are provided in the DTO.
//
// Parameters:
//   - article (*models.Article): The existing article model to update
//   - updateArticleDto (dto.UpdateArticle): The DTO containing update data
//
// Returns:
//   - *models.Article: The updated Article model
func updateArticleFromDto(article *models.Article, updateArticleDto dto.UpdateArticle) *models.Article {
	if updateArticleDto.Title != "" && updateArticleDto.Title != article.Title {
		article.Title = updateArticleDto.Title
	}

	if updateArticleDto.Slug != "" && updateArticleDto.Slug != article.Slug {
		article.Slug = updateArticleDto.Slug
	}

	if updateArticleDto.Content != "" && updateArticleDto.Content != article.Content {
		article.Content = updateArticleDto.Content
	}

	if updateArticleDto.Excerpt != "" && updateArticleDto.Excerpt != article.Excerpt.String {
		article.Excerpt = dbNull.String(updateArticleDto.Excerpt)
	}

	if updateArticleDto.CoverImage != "" && updateArticleDto.CoverImage != article.CoverImage.String {
		article.CoverImage = dbNull.String(updateArticleDto.CoverImage)
	}

	if updateArticleDto.SEODescription != "" && updateArticleDto.SEODescription != article.SEODescription.String {
		article.SEODescription = dbNull.String(updateArticleDto.SEODescription)
	}

	if updateArticleDto.SEOKeywords != "" && updateArticleDto.SEOKeywords != article.SEOKeywords.String {
		article.SEOKeywords = dbNull.String(updateArticleDto.SEOKeywords)
	}

	if updateArticleDto.YouTubeURL != "" && updateArticleDto.YouTubeURL != article.YouTubeURL.String {
		article.YouTubeURL = dbNull.String(updateArticleDto.YouTubeURL)
	}

	if updateArticleDto.TikTokURL != "" && updateArticleDto.TikTokURL != article.TikTokURL.String {
		article.TikTokURL = dbNull.String(updateArticleDto.TikTokURL)
	}

	// Handle status change
	if updateArticleDto.Status != "" && updateArticleDto.Status != article.Status {
		article.Status = updateArticleDto.Status

		// Update published_at if status is changed to published
		if article.Status == types.ArticleStatusPublished && !article.PublishedAt.Valid {
			article.PublishedAt = dbNull.Time(time.Now())
		}
	}

	article.UpdatedAt = dbNull.Time(time.Now())

	return article
}
