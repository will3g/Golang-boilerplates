package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Article struct {
	db       *sql.DB
	ID       string
	Title    string
	Subtitle string
	Body     string
	AuthorID string
}

func ArticleInstance(db *sql.DB) *Article {
	return &Article{db: db}
}

func (articlePointer *Article) Create(title, subtitle, body, authorID string) (*Article, error) {
	id := uuid.New().String()
	_, err := articlePointer.db.Exec("INSERT INTO articles (id, title, subtitle, body, author_id) VALUES ($1, $2, $3, $4, $5)", id, title, subtitle, body, authorID)
	if err != nil {
		return nil, err
	}
	return &Article{
		ID:       id,
		Title:    title,
		Subtitle: subtitle,
		Body:     body,
		AuthorID: authorID,
	}, nil
}

func (articlePointer *Article) FindAll() ([]Article, error) {
	rows, err := articlePointer.db.Query("SELECT id, title, subtitle, body, author_id FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []Article{}

	for rows.Next() {
		var id, title, subtitle, body, authorID string
		if err := rows.Scan(&id, &title, &subtitle, &body, &authorID); err != nil {
			return nil, err
		}
		articles = append(articles, Article{
			ID:       id,
			Title:    title,
			Subtitle: subtitle,
			Body:     body,
			AuthorID: authorID,
		})
	}

	return articles, nil
}

func (articlePointer *Article) FindByAuthorID(authorID string) ([]Article, error) {
	rows, err := articlePointer.db.Query("SELECT id, title, subtitle, body, author_id FROM courses WHERE author_id = $1", authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []Article{}

	for rows.Next() {
		var id, title, subtitle, body, authorID string
		if err := rows.Scan(&id, &title, &subtitle, &body, &authorID); err != nil {
			return nil, err
		}
		articles = append(articles, Article{
			ID:       id,
			Title:    title,
			Subtitle: subtitle,
			Body:     body,
			AuthorID: authorID,
		})
	}

	return articles, nil
}
