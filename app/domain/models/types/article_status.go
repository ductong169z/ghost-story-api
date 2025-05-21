package types

// ====================================================================
// ============================ Data Types ============================
// ====================================================================

type ArticleStatus string

// Article status types
const (
	ArticleStatusDraft     ArticleStatus = "draft"
	ArticleStatusPublished ArticleStatus = "published"
	ArticleStatusArchived  ArticleStatus = "archived"
)

var ArticleStatusList = []ArticleStatus{
	ArticleStatusDraft,
	ArticleStatusPublished,
	ArticleStatusArchived,
}

// ====================================================================
// ============================= Methods ==============================
// ====================================================================

type articleStatusCollection []ArticleStatus

// String converts a collection of ArticleStatus to an array of strings.
//
// Returns:
//   - []string: Array of ArticleStatus values as strings (e.g. ["draft", "published", "archived"])
func (e articleStatusCollection) String() []string {
	result := make([]string, len(e))
	for i, v := range e {
		result[i] = string(v)
	}

	return result
}

// ArticleStatusArrStr converts variable number of ArticleStatus to array of strings.
//
// Parameters:
//   - articleStatus: Variable number of ArticleStatus values
//
// Returns:
//   - []string: Array of ArticleStatus values as strings (e.g. ["draft", "published", "archived"])
func ArticleStatusArrStr(articleStatus ...ArticleStatus) []string {
	return append(articleStatusCollection{}, articleStatus...).String()
}
