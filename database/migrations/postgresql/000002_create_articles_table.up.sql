CREATE TYPE article_status AS ENUM ('draft', 'published', 'archived');

CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    excerpt TEXT,
    content TEXT NOT NULL,
    cover_image VARCHAR(255),
    status article_status DEFAULT 'draft',
    seo_description VARCHAR(300),
    seo_keywords TEXT,
    author_id INT NOT NULL,
    published_at TIMESTAMP,
    youtube_url VARCHAR(255),
    tiktok_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT fk_articles_author
        FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_articles_status ON articles(status);
CREATE UNIQUE INDEX idx_articles_slug ON articles(slug);
