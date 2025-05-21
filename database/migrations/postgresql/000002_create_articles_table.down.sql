-- Drop indexes
DROP INDEX IF EXISTS idx_articles_status;
DROP INDEX IF EXISTS idx_articles_author;
DROP INDEX IF EXISTS idx_articles_slug;

-- Drop table
DROP TABLE IF EXISTS articles;
