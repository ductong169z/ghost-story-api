-- Add view_count column to articles table with default value 0
ALTER TABLE articles ADD COLUMN view_count INT NOT NULL DEFAULT 0;

-- Create index for view_count to optimize queries that sort by view count
CREATE INDEX idx_articles_view_count ON articles(view_count);
