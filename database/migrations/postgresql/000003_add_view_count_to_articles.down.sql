-- Drop the index first
DROP INDEX IF EXISTS idx_articles_view_count;

-- Remove the view_count column from articles table
ALTER TABLE articles DROP COLUMN IF EXISTS view_count;
